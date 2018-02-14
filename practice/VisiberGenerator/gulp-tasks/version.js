var gulp = require('gulp');
var RevAll = require('gulp-rev');
var es = require('event-stream');
var revreplace = require('gulp-rev-replace');

gulp.task('copy-build', ['build-styling', 'build-js', 'build-static'], function () {
    console.log("Copy build/config.json to dist/");
    var copyConfig = gulp.src('./resources/app/config.json')
        .pipe(gulp.dest('./gulp/dist/resources/app/'));

    console.log("Copy build/index.html to dist/");
    var copyIndex = gulp.src('./resources/app/index.html')
        .pipe(gulp.dest('./gulp/dist/resources/app/'));

    console.log("Copy assets");
    var copyAssets = gulp.src('./resources/app/assets/*/**')
        .pipe(gulp.dest('./gulp/dist/resources/app/assets/'));

    console.log("Copy bower_components");
    var copyBowers = gulp.src('./resources/app/bower_components/*/**')
        .pipe(gulp.dest('./gulp/dist/resources/app/bower_components/'));

    console.log("Copy bundler.json");
    var copyBundle = gulp.src('./resources/app/bundler.json')
        .pipe(gulp.dest('./gulp/dist/'));

    console.log("Copy Go Files");
    var copyGo = gulp.src(['./resources/app/*.go'])
        .pipe(gulp.dest('./gulp/dist/'));

    console.log("Copy XML Files");
    var copyXML = gulp.src(['./resources/app/*.xml'])
        .pipe(gulp.dest('./gulp/dist/resources/app/'));

    console.log("Copy Icon");
    var copyIcon = gulp.src(['./resources/icon.icns', './resources/icon.ico', './resources/icon.png'])
        .pipe(gulp.dest('./gulp/dist/resources/'));

    return es.merge([copyAssets, copyBowers, copyConfig, copyIndex, copyBundle, copyGo, copyXML, copyIcon])
});

gulp.task('version', ['copy-build'], function () {
    console.log("Versioning /css");
    var versionCSS = gulp.src('./gulp/build/resources/app/css/*')
        .pipe(RevAll())
        .pipe(gulp.dest('./gulp/dist/resources/app/css/'))
        .pipe(RevAll.manifest())
        .pipe(gulp.dest('./gulp/dist/resources/app/css/'));

    console.log("Versioning /js");
    var versionJS = gulp.src('./gulp/build/resources/app/js/*')
        .pipe(RevAll())
        .pipe(gulp.dest('./gulp/dist/resources/app/js/'))
        .pipe(RevAll.manifest())
        .pipe(gulp.dest('./gulp/dist/resources/app/js/'));

    return es.merge([versionCSS, versionJS])
});

//get all images from all directory and add them into gulp/dist/img
gulp.task('deploy-index', ['version'], function () {
    console.log("Replacing version in index.html");
    var replaceIndex = gulp.src('./gulp/dist/resources/app/index.html')
        .pipe(revreplace({manifest: gulp.src('./gulp/dist/resources/app/js/rev-manifest.json')}))
        .pipe(revreplace({manifest: gulp.src('./gulp/dist/resources/app/css/rev-manifest.json')}))
        .pipe(gulp.dest('./gulp/dist/resources/app/'));

    return es.merge([replaceIndex])
});
