const Generator = require('yeoman-generator');
const ejs = require('ejs');

const template = "\n\
GET     /<%- entity.lname %>                       <%- entity.name %>.Index\n\
GET     /<%- entity.lname %>/create                <%- entity.name %>.Add\n\
POST    /<%- entity.lname %>/create                <%- entity.name %>.AddPost\n\
GET     /<%- entity.lname %>/:id                   <%- entity.name %>.Detail\n\
GET     /<%- entity.lname %>/:id/edit              <%- entity.name %>.Edit\n\
POST    /<%- entity.lname %>/:id/edit              <%- entity.name %>.Save\n\
DELETE  /<%- entity.lname %>/:id                   <%- entity.name %>.Delete\n\
";

module.exports = class extends Generator {
    writing(){
        this.routerTxt = "";
        this.options.entities.forEach(entity => {
            this.routerTxt += ejs.render(template, {entity: entity});
            this.fs.copyTpl(
                this.templatePath("sitemap.conf"),
                this.destinationPath("conf/sitemap/"+entity.lname+".conf"),
                {entity: entity})
        });

        var routes = this.fs.read(this.destinationPath('conf/routes'), {defaults: ""});
        routes += this.routerTxt;
        this.fs.write(this.destinationPath('conf/routes'), routes);
    }
};