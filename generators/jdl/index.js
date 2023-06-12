const Generator = require('yeoman-generator');
const jdlCore = require('jhipster-core');

module.exports = class extends Generator {

    data = null;

    constructor(args, opts) {
        super(args, opts);
        this.argument("jdlfile", { type: String, required: true });
    }

    'initializing' () {
        this.data = jdlCore.parseFromFiles([this.options.jdlfile]);
        this.composeWith('gorilla:model', this.data);
        this.composeWith('gorilla:controller', this.data);
        this.composeWith('gorilla:router', this.data);
        this.composeWith('gorilla:service', this.data);
    }

    prompting() {
        this.options.config = this.config.getAll();
    }
    
  };