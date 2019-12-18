/*
Author       : theme_crazy
Template Name: Cryptoma - Bitcoin & Cryptocurrency ICO Landing Page
Version      : 1.0
*/

(function($)
{
	"use strict";
	
	// Preloader
	jQuery(window).on('load', function() {
		preloader();
	});
	
	var headerHeight = jQuery('.navbar').outerHeight();
	jQuery('.navbar-nav li a').on('click', function(event) {
		jQuery('.navbar-nav li').removeClass('active');
		jQuery(this).parent().addClass('active');
		var $anchor = jQuery(this);
		
		jQuery('html, body').stop().animate({
			scrollTop: jQuery($anchor.attr('href')).offset().top-headerHeight
		}, 1000, 'easeInOutExpo');
		event.preventDefault();
	});
	
	jQuery(".navbar-nav li a").on("click",function(event){
		jQuery(".navbar-collapse").removeClass('show');
		jQuery('.navbar-toggler').addClass('collapsed');
	});
	
	jQuery('.tilt-img img').tilt({
		maxTilt:6					
	});
	
	// Animation section
	 if(jQuery('.wow').length){
		var wow = new WOW(
		  {
			boxClass:     'wow',      // animated element css class (default is wow)
			animateClass: 'animated', // animation css class (default is animated)
			offset:       0,          // distance to the element when triggering the animation (default is 0)
			mobile:       true,       // trigger animations on mobile devices (default is true)
			live:         true       // act on asynchronously loaded content (default is true)
		  }
		);
		wow.init();
	}
	
	// Back to top 		
	jQuery('.back-top a').on('click', function(event) {
		jQuery('body,html').animate({scrollTop:0},800);
		return false;
	});
	
	jQuery(window).on('scroll', function() {
		
		// Back to top 
		if(jQuery(this).scrollTop()>150){
			jQuery('.back-top').fadeIn();
		}
		else{
			jQuery('.back-top').fadeOut();
		}
	});
	
	// Preload
	function preloader(){
		jQuery(".preloaderimg").fadeOut();
		jQuery(".preloader").delay(200).fadeOut("slow").delay(200, function(){
			jQuery(this).remove();
		});
	}
	
	
})(jQuery);	