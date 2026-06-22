"use strict";

const cp = require("child_process");
const fs = require("fs");

// eval with a non-literal expression — no-eval / security/detect-eval-with-expression
function evaluate(input) {
  return eval(input);
}

// child_process.exec with user input — security/detect-child-process
function runCommand(userCmd) {
  return cp.exec(userCmd);
}

// reading a non-literal filename — security/detect-non-literal-fs-filename
function readUserFile(name) {
  return fs.readFileSync(name);
}

// bracket access with a user-controlled key — security/detect-object-injection
function getProp(obj, key) {
  return obj[key];
}

module.exports = { evaluate, runCommand, readUserFile, getProp };
