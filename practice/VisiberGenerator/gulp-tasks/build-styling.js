var gulp = require('gulp');
var concat = require('gulp-concat');
var minifyCss = require('gulp-minify-css');
var es = require('event-stream');

var cssFiles = [
    'bower_components/bootstrap/dist/css/bootstrap.css'
];

//call this task will build all styling files above
gulp.task('build-styling', function() {
    console.log("Building vendor.css");
    var vendorBuild = gulp.src(cssFiles)
        // .pipe(concat('vendor.css'))
        .pipe(minifyCss({"mangle": false}))
        .pipe(gulp.dest('./gulp/build/resources/app/css/'));

    console.log("Build app.css");
    var appBuild = gulp.src('./resources/app/app.css')
        .pipe(concat('app.css'))
        .pipe(minifyCss())
        .pipe(gulp.dest('./gulp/build/resources/app/css/'));

    return es.merge([vendorBuild, appBuild])
});