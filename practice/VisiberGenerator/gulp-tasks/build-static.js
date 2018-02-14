var gulp = require('gulp');
var es = require('event-stream');

gulp.task('build-static', function () {
    console.log("Build config.json");
    var copyConfig = gulp.src('./resources/app/config.json')
        .pipe(gulp.dest('./gulp/build/resources/app/'));

    console.log("Build index.html");
    var copyIndex = gulp.src('./resources/app/index.html')
        .pipe(gulp.dest('./gulp/build/resources/app/'));

    return es.merge([copyConfig, copyIndex])
});

