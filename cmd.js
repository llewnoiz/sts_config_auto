const { Observable } = require('rxjs');
const { spawn } = require('child_process');

const format = (type, name, status, description, errCode) => ({
  type,
  name,
  status,
  description,
  errCode,
});

const Command$ = (command, subCommand, job) => new Observable((subscriber) => {

  const proc = spawn(command, job.args);
  let buf = null;

  proc.stdout?.on('data', (data) => {
    buf = data?.toString();
    subscriber.next(format(command, subCommand, 'process', buf, proc.exitCode));
  });

  proc.stderr?.on('data', (data) => {
    buf = data?.toString();
    subscriber.next(format(
      command,
      subCommand,
      'process',
      buf,
      proc.exitCode,
    ));
  });

  proc.on('error', (data) => {
    buf = data?.toString();
    subscriber.error(format(
      command,
      subCommand,
      'process',
      buf,
      proc.exitCode,
    ));
  });
  proc.on('close', async (data) => {
    // eslint-disable-next-line no-unused-expressions
    if (proc.exitCode === 0) {

      subscriber.next(
        format(command, subCommand, 'process', buf, proc.exitCode),
      );

      subscriber.complete();
    } else {
      subscriber.error(
        format(command, subCommand, 'process', buf, proc.exitCode),
      );
      subscriber.unsubscribe();
    }
  });

});

export default { Command$ };