var requireDir = require('require-dir');
var dir = requireDir('./gulp-tasks');

//******************************************************************************************************
// gulp tasks run sequence
// 1) build-styling (build-styling.js)
// 	    - build gulp/build/css/vendor.css from cssFiles
// 	    - build gulp/build/css/components.css from app/components/**/*.css
// 	    - build gulp/build/css/portal.css from app/portal/portal.css
// 	    - build gulp/build/css/app.css from app/app.css
// 	    - build gulp/dist/css/fonts/ from list
// 	    - build gulp/dist/fonts/ from list
// 	    - build gulp/dist/img/ from app/**/* jpg, png, gif
// 2) build-js (build-js.js)
// 	    - build gulp/build/js/IE.js from list
// 	    - build gulp/build/js/vendor.js from jsFiles
// 	    - build gulp/build/js/components.js from app/components/**/*.js
// 	    - build gulp/build/js/portal.js from app/portal/**/*.js
// 	    - build gulp/build/templates/templates.js from app/**/*.html exclude build and dist folder
// 	    - build gulp/build/app/app.js from app/app.js
// 3) build-static (build-static.js)
// 	    - copy app/index.html to gulp/build/index.html
// 	    - copy app/config.jsonl to gulp/build/config.json
// 4) version (version.js)
//      - copy gulp/build/index.html to gulp/dist/
//      - copy gulp/build/config.json to gulp/dist/
//      - rev gulp/build/css to gulp/dis/css/
//      - rev gulp/build/js to gulp/dis/js/
//      - rev gulp/build/templates to gulp/dist/templates/css/
//      - rev gulp/build/app to gulp/dist/app/
// 5) deploy-index (deploy-index.js)
// 	    - rev replace directory app, js, css and templates file names in dist/index.html
// 6) build (browser-sync.js)
// 	    - sync reload page
// 7) run-local (run-local.js)
// 	    - run gulp/dist/
//
// deploy dev (deploy-release.js) need to run 'deploy-index' (deploy-index.js) first if there are changes to project
//      - deploy release build
//      - deploy release
//      - deploy s3
//
//******************************************************************************************************