@set @junk=1 /*
@cscript //nologo //e:jscript "%~f0" %* & @ping -n 10 localhost > nul & @goto :eof
*/
function base() {
  WScript.Echo("Hello from lib/base");
}

function foo() {
  WScript.Echo("Hello from lib/foo");
}
console.log('From boo/foo/bar.js');
console.log('���{��e�X�g');

console.log('From main/main/hoge.js');

function main(arg) {
  console.log("main�֐��ł��B");
}

function test() {
  WScript.Echo("Hello from test");
}
