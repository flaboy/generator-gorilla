const Generator = require('yeoman-generator');

function upCaseFirst(str){
  return str.charAt(0).toUpperCase() + str.slice(1);
}

function lowCaseFirst(str){
  return str.charAt(0).toLowerCase() + str.slice(1);
}

function camelToHyphen(str) {
  str = str.replace(/([A-Z])/g, '_$1').toLowerCase();
  if(str.charAt(0) === '_'){
    str = str.slice(1);
  }
  return str;
}

function typeToGoType(str, entity) {
  switch(str){
    case 'String':
      return 'string';
    case 'Integer':
      return 'int';
    case 'Long':
      return 'int64';
    case 'Float':
      return 'float32';
    case 'Double':
      return 'float64';
    case 'BigDecimal':
      return 'float64';
    case 'LocalDate':
      entity.imports['time'] = true;
      return 'time.Time';
    case 'Instant':
      entity.imports['time'] = true;
      return 'time.Time';
    case 'ZonedDateTime':
      entity.imports['time'] = true;
      return 'time.Time';
    case 'Boolean':
      return 'bool';
    default:
      return str;
  }
}

module.exports = class extends Generator {

    prompting() {
      this.options.config = this.config.getAll();
    }

    prepareEnums() {
      this.enums = {};
      this.options.enums.forEach(enumType => {
        this.enums[enumType.name] = enumType;
      });
    }

    prepareEntities() {
      this.entities = {};
      this.options.entities.forEach(entity => {
        this.entities[entity.name] = entity;
        entity.enums = [];
        entity.imports = {};
        entity.body.forEach(field => {
          field.Uname = upCaseFirst(field.name);
          field.gormOptions = "";
          if(this.enums[field.type]){
            field.isEnum = true;
            if(!this.enums[field.type].appened){
              entity.enums.push(this.enums[field.type]);
              this.enums[field.type].appened = true;
              // entity.gormType = 'string';
            }
          }else{
            field.type = typeToGoType(field.type, entity);
          }
          field.columnName = camelToHyphen(field.name);
        });
      });
    }

    prepareRelationShips() {
      this.options.relationships.forEach(relationship => {
        if(relationship.cardinality === 'OneToOne'){
          this.entities[relationship.to.name].body.push({
            name: relationship.from.name,
            type: upCaseFirst(relationship.from.name),
            Uname: upCaseFirst(relationship.from.name),
            skipInForm: true
          });
          this.entities[relationship.to.name].body.push({
            name: relationship.from.name+"ID",
            type: "int64",
            Uname: upCaseFirst(relationship.from.name+"ID"),
            columnName: camelToHyphen(relationship.from.name+"_id")
          });
          this.entities[relationship.from.name].body.push({
            name: relationship.to.name,
            type: upCaseFirst(relationship.to.name),
            Uname: upCaseFirst(relationship.to.name),
            skipInForm: true
          });
        }else if(relationship.cardinality === 'OneToMany'){
          this.entities[relationship.to.name].body.push({
            comment: true,
            name: relationship.from.name,
            type: upCaseFirst(relationship.from.name),
            Uname: upCaseFirst(relationship.from.name),
            skipInForm: true
          });
          this.entities[relationship.to.name].body.push({
            name: relationship.from.name+"ID",
            type: "int64",
            Uname: upCaseFirst(relationship.from.name+"ID"),
            columnName: camelToHyphen(relationship.from.name+"_id")
          });
        }else if(relationship.cardinality === 'ManyToOne'){
          this.entities[relationship.from.name].body.push({
            comment: true,
            name: relationship.to.name,
            type: upCaseFirst(relationship.to.name),
            Uname: upCaseFirst(relationship.to.name),
            skipInForm: true
          });
          this.entities[relationship.from.name].body.push({
            name: relationship.to.name+"ID",
            type: "int64",
            Uname: upCaseFirst(relationship.to.name+"ID"),
            columnName: camelToHyphen(relationship.to.name+"_id")
          });
        }else if(relationship.cardinality === 'ManyToMany'){
          this.entities[relationship.from.name].body.push({
            name: relationship.to.name,
            type: "[]*"+upCaseFirst(relationship.to.name),
            Uname: upCaseFirst(relationship.to.name),
            skipInForm: true
          });
          this.entities[relationship.to.name].body.push({
            name: relationship.from.name,
            type: "[]*"+upCaseFirst(relationship.from.name),
            Uname: upCaseFirst(relationship.from.name),
            skipInForm: true
          });
        }
      });
    }
    
    writing(){
      this.options.entities.forEach(entity => {
        this.fs.copyTpl(
            this.templatePath('model.go'),
            this.destinationPath('app/models/'+lowCaseFirst(entity.name)+'.go'),
            {config:this.options.config, entity:entity}
          );
        });
      }
  };