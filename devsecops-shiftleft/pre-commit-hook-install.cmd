@ECHO OFF
FOR /f %%i IN ('git rev-parse --git-dir') DO SET GIT_HOOKS_DIR=%%i/hooks
echo %GIT_HOOKS_DIR%
FOR /f %%i IN ('git rev-parse --show-toplevel') DO SET GIT_HOOKS_INSTALL=%%i/devsecops-shiftleft
echo %GIT_HOOKS_INSTALL%
echo "%GIT_HOOKS_INSTALL:/=\%\pre-commit-hook"
COPY "%GIT_HOOKS_INSTALL:/=\%\pre-commit-hook" "%GIT_HOOKS_DIR%/pre-commit"