var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var es = require('event-stream');

var jsFiles = [
    'bower_components/jquery/dist/jquery.js',
    'bower_components/bootstrap/dist/js/bootstrap.js',
    'bower_components/angular/angular.js',
    'bower_components/angular-bootstrap/ui-bootstrap-tpls.js'
];

// call this will build all tasks above
gulp.task('build-js', function() {
    console.log("Building vendor.js");
    var vendorBuild = gulp.src(jsFiles)
        // .pipe(concat('vendor.js'))
        .pipe(uglify({"mangle": false}))
        .pipe(gulp.dest('./gulp/build/resources/app/js/'));

    console.log("Building app.js");
    //*** Compile App JS
    var appBuild = gulp.src('./resources/app/app.js')
        .pipe(concat('app.js'))
        .pipe(uglify({"mangle": false}))
        .pipe(gulp.dest('./gulp/build/resources/app/js/'));

    return es.merge([vendorBuild, appBuild])
});