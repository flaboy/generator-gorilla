const Generator = require('yeoman-generator');
const ejs = require('ejs');

const template = "\n\
GET     /<%- entity.lname %>                       <%- entity.name %>Api.Index\n\
POST    /<%- entity.lname %>                       <%- entity.name %>Api.Add\n\
GET     /<%- entity.lname %>/:id                   <%- entity.name %>Api.Detail\n\
PUT     /<%- entity.lname %>/:id                   <%- entity.name %>Api.Save\n\
DELETE  /<%- entity.lname %>/:id                   <%- entity.name %>Api.Delete\n\
";

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

    writing(){
      this.routerTxt = "";      
      this.options.entities.forEach(entity => {
        this.routerTxt += ejs.render(template, {entity: entity});        
        entity.lname = lowCaseFirst(entity.name)
          this.fs.copyTpl(
            this.templatePath('controller.go'),
            this.destinationPath('app/api/app/controllers/'+entity.lname+'Api.go'),
            {config:this.options.config, entity:entity}
          );
        });

        var routes = this.fs.read(this.destinationPath('app/api/conf/routes'), {defaults: ""});
        routes += this.routerTxt;
        this.fs.write(this.destinationPath('app/api/conf/routes'), routes);
      } 
  };