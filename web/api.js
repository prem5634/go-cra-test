"use strict";

const cp = require("child_process");

// eval with a non-literal expression — no-eval / security/detect-eval-with-expression
function compute(input) {
  return eval(input);
}

// child_process.exec with user input — security/detect-child-process
function shell(userCmd) {
  return cp.exec(userCmd);
}

// bracket access with a user-controlled key — security/detect-object-injection
function pick(obj, key) {
  return obj[key];
}

module.exports = { compute, shell, pick };
