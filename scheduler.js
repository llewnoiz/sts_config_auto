class Scheduler {
    constructor(recipe) {

      this.tasks = [];
      this.recipe = recipe;
      this.currentTask = null;
    }
  
    add(task$) {
      this.tasks.push(task$);
    }
  
    process() {
      const self = this;
      const data = self.recipe;
      this.currentTask = this.tasks.shift();
      try {
        this.currentTask.task$?.subscribe({
          next(message) {
            data.phase = message.type;
            data.description = message?.description === undefined ? '' : message?.description;
            if (message.errCode === 0) {
              

            } else if (message.errCode === null) {
              
              
            } else {

            }
          },
          error(err) {
            console.error(err);
        
          },
          complete() {
            if (!self.tasks.length) {
                console.log('complete');
            } else {
              self.process();
            }
          },
        });
      } catch (err) {
        console.error(err);
      }
  
      return 0;
    }
  }

  module.exports = { Scheduler };