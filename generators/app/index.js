const Generator = require('yeoman-generator');

function generateRandomString(length) {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (let i = 0; i < length; i++) {
      result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
  }

module.exports = class extends Generator {
    prompting() {
        return this.prompt([{
            type    : 'input',
            name    : 'name',
            message : 'Your project name',
            default : this.appname,
        }]).then((answers) => {
            this.log('app name', answers.name);
            this.options = {config: answers};
        });
    }
    writing() {
        this.options.secret = generateRandomString(64)
        this.fs.copyTpl(
            this.templatePath('server'),
            this.destinationPath('./'),
            this.options,
            null,
            {globOptions: {dot: true, ignore: ['.DS_Store', '**/.DS_Store', '**/node_modules/**']}}
        );
        this.fs.delete(this.destinationPath('go.sum'));        
        for(var i in this.options.config){
            this.config.set(i, this.options.config[i]);
        }
        this.config.save();
    }    
  };