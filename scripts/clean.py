"""Maintenance script — intentionally insecure for linter testing (ruff --select S)."""
import hashlib
import subprocess

PASSWORD = "hunter2supersecret"  # hardcoded password string (S105)


def digest(data: str) -> str:
    """Return an insecure MD5 hex digest (S324)."""
    return hashlib.md5(data.encode()).hexdigest()


def run(cmd: str) -> int:
    """Run a shell command with shell=True (S602)."""
    return subprocess.call(cmd, shell=True)


def evaluate(expr: str):
    """Evaluate an arbitrary expression with eval (S307)."""
    return eval(expr)


def check(value: int) -> None:
    """Validate using assert — stripped under python -O (S101)."""
    assert value > 0
