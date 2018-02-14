var gulp = require('gulp');
var revreplace = require('gulp-rev-replace');
var es = require('event-stream');
var exec = require('gulp-exec');

//get all images from all directory and add them into gulp/dist/img
gulp.task('revision-index', ['version'], function () {
    console.log("Replacing version in index.html");
    var replaceIndex = gulp.src('./gulp/build/resources/app/index.html')
        .pipe(revreplace({manifest: gulp.src('./gulp/dist/resources/app/js/rev-manifest.json')}))
        .pipe(revreplace({manifest: gulp.src('./gulp/dist/resources/app/css/rev-manifest.json')}))
        .pipe(gulp.dest('./gulp/dist/resources/app/'));

    return es.merge([replaceIndex])
});

gulp.task('build-gui', ['revision-index'], function (cb) {
    exec('sh build.sh', function (err, stdout, stderr) {
        console.log(stdout);
        console.log(stderr);
        cb(err);
    });
});