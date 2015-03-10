var gulp  = require('gulp');
var clean = require('del');
var gulpif = require('gulp-if');
var uglify = require('gulp-uglify');
var mincss = require('gulp-minify-css');
var useref = require('gulp-useref');
var bdiff  = require('gulp-bytediff');

gulp.task('clean', function() {
	clean(['public']);
});

gulp.task('default', function() {
	// Copy templates
	gulp.src('site/templates/**/*.html')
		.pipe(gulp.dest('public/templates'))

	var assets = useref.assets();
	return gulp.src('site/index.html')
		.pipe(assets)
		.pipe(bdiff.start())
		.pipe(gulpif('*.js', uglify()))
		.pipe(gulpif('*.css', mincss()))
		.pipe(bdiff.stop())
		.pipe(assets.restore())
		.pipe(useref())
		.pipe(gulp.dest('public'))
});