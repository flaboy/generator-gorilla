const Generator = require('yeoman-generator');

function upCaseFirst(str){
  return str.charAt(0).toUpperCase() + str.slice(1);
}

function lowCaseFirst(str){
  return str.charAt(0).toLowerCase() + str.slice(1);
}

function camelToHyphen(str) {
  return str.replace(/([A-Z])/g, '_$1').toLowerCase();
}

module.exports = class extends Generator {
  
    prompting() {
      this.options.config = this.config.getAll();
    }

    prepareServices(){
      this.options.options.service.forEach(service => {
        console.info("service",service);
      });
    }

    prepareEntities() {
      this.options.entities.forEach(entity => {
        entity.body.forEach(field => {
          if(field.type!="string" && !field.skipInForm){
            switch(field.type){
              case "int":
                field.typeChanger = "toInt";
                break;
              case "int64":
                field.typeChanger = "toInt64";
                break;
              case "float32":
                field.typeChanger = "toFloat32";
                break;
              case "float64":
                field.typeChanger = "toFloat64";
                break;
              case "bool":
                field.typeChanger = "toBool";
                break;
              case "time.Time":
                field.typeChanger = "toTime";
                break;
              default:
                if(field.isEnum){
                  field.typeChanger = "models."+field.type;
                }
            }
            entity.useStrconv = true;
          }
        });
      })
    }

    writing(){
      // this.options.entities.forEach(entity => {
      //   entity.lname = lowCaseFirst(entity.name)
      //     this.fs.copyTpl(
      //       this.templatePath('controller.go'),
      //       this.destinationPath('app/api/app/controllers/'+entity.lname+'.go'),
      //       {config:this.options.config, entity:entity}
      //     );
      //   });
      // } 
  };