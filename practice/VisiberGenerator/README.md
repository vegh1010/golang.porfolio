This steps represent how you can compile the project and create the GUI interface

# Step 1: install the bundler

Run the following command:

    $ go get -u github.com/asticode/go-astilectron-bundler/...
    
And don't forget to add `$GOPATH/bin` to your `$PATH`.
    
# Step 2: install all components

Run the following commands:

    $ npm install
    $ bower install
      
# Step 3: run gulp which will build the app in gulp folder

Run the following commands:

    $ gulp build-gui
      
# Step 4: run this to build the app GUI

Run the following commands:

    $ sh build-gui
    
# Step 5: test the app

The result is in the `gulp/dist/output/<your os>-<your arch>` folder and is waiting for you to test it!

# Step 6: bundle the app for more environments

To bundle the app for more environments, add an `environments` key to the bundler configuration (`bundler.json`):

```json
"environments": [
  {"arch": "amd64", "os": "linux"},
  {"arch": "386", "os": "windows"}
]
```

# Step 7: configure api call

To configure api call, modify config.json in resources/app/config.json

and repeat **step 3**.