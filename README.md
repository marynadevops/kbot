# kbot

devops application from scratch


# devsecops



## DevSecOps - shift left

Pre-commit hook.

To install

```terminal
"./devsecops-shiftleft/pre-commit-hook-install"
```

```sh
chmod +x ./devsecops-shiftleft/pre-commit-hook-install
./devsecops-shiftleft/pre-commit-hook-install
```

```cmd
.\devsecops-shiftleft\pre-commit-hook-install
```

```
git commit -m 1 || git status
```

```terminal
git config devsecops.gitleaks.protect.enabled false
```

```sh
GIT_HOOKS_DIR=$(git rev-parse --git-dir)/hooks
cp $(git rev-parse --show-toplevel)/devsecops-shiftleft/pre-commit-hook "$GIT_HOOKS_DIR/pre-commit"
chmod +x "$GIT_HOOKS_DIR/pre-commit"
```

```cmd
GIT_HOOKS_DIR=$(git rev-parse --git-dir)/hooks
cp $(git rev-parse --show-toplevel)/devsecops-shiftleft/pre-commit-hook "$GIT_HOOKS_DIR/pre-commit"
chmod +x "$GIT_HOOKS_DIR/pre-commit"

@ECHO OFF
FOR /f %%i IN ('git rev-parse --git-dir') DO SET GIT_HOOKS_DIR=%%i/hooks
echo %GIT_HOOKS_DIR%
FOR /f %%i IN ('git rev-parse --show-toplevel') DO SET GIT_HOOKS_INSTALL=%%i/devsecops-shiftleft
echo %GIT_HOOKS_INSTALL%
echo "%GIT_HOOKS_INSTALL%/pre-commit-hook:/=\"
COPY "%GIT_HOOKS_INSTALL%/pre-commit-hook:/=\" "%GIT_HOOKS_DIR%/pre-commit"

cp $(git rev-parse --show-toplevel)/devsecops-shiftleft/pre-commit-hook "$GIT_HOOKS_DIR/pre-commit"

devsecops-shiftleft\pre-commit-hook-install
```

<details>
  <summary>Click me</summary>

### Heading

  1. Foo
  2. Bar
     * Baz
     * Qux

### Some Code

  ```js
  function logSomething(something) {
    console.log('Something', something);
  }
  ```

</details>

<table>
  <tr>
    <th>Column 1</th>
    <th>Column 2</th>
    <th>Column 3</th>
  </tr>
  <tr>
    <td>Linux</td>
    <td>MacOS</td>
    <td>Windows</td>
  </tr>
  <tr>
    <td>Text</td>
    <td>
      <details>
        <summary>Click here for MacOS</summary>
          <pre><code># code
This is a multiline
code fragment
          </code></pre>
      </details>
    </td>
    <td>
      <details>
        <summary>Click here for Windows/Git Bash/MINGW</summary>
          <pre><code>
  This is a multiline
  code fragment
          </code></pre>
      </details>
    </td>
  </tr>
</table>

<style>
/** CSS styles go here */
.tabs {
  display: flex;
}

.tab-content {
  display: none;
}

input[type="radio"]:checked + label + .tab-content {
  display: block;
}
</style>

<!-- # Tabbed Layout Example -->

<div class="tabs">
  <input type="radio" name="tab" id="tab1" checked>
  <label for="tab1">Tab 1</label>
  <div class="tab-content">
    Content for Tab 1 goes here.
  </div>

  <input type="radio" name="tab" id="tab2">
  <label for="tab2">Tab 2</label>
  <div class="tab-content">
    <details>
      <summary>Click me</summary>

### Heading1

      1. Foo
      2. Bar
        * Baz
        * Qux

    ### Some Code

      ```js
      function logSomething(something) {
        console.log('Something', something);
      }
      ```

    </details>

  </div>

  <input type="radio" name="tab" id="tab3">
  <label for="tab3">Tab 3</label>
  <div class="tab-content">
    Content for Tab 3 goes here.
  </div>
</div>

