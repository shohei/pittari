/*global module:false*/
module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    watch: {
      scad: {
        files: '*.scad',
        tasks: ['build']
      },
    build:{
        
    }
  });

  // These plugins provide necessary tasks.
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task.
  // grunt.registerTask('default', ['jshint', 'nodeunit', 'concat', 'uglify']);

};
