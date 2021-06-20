const exec = require('child_process').exec;
const spawn = require('child_process').spawn;

// child_process is for short running operations
exec('git version', (err, stdout) => {
    if (err) throw err;

    console.log(stdout.trim());
});

const childProcess = spawn('node', ['longRunningProgram']);
childProcess.stdout.on('data', (data) => {
    console.log(`STDOUT: ${data.toString()}`);
});

childProcess.on('close', () => {
    console.log('Child process has ended');
    process.exit();
});

setTimeout(() => {
    // child process is prepared to exit when it receives data
    childProcess.stdin.write('stopit');
}, 4000);

