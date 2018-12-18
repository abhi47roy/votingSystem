var playPauseStatus = 0;
var animationId = 0;


var isNotPaused = true;


//call this function from vote button (modify onload thing)
function startAnimation() {

  retrieveTransactionHash();
  toggleImage();
  $( "#togglePlayPause" ).on( "click", function() {
  	playPauseExecute();
	});
};



function playPauseExecute() {
	if(playPauseStatus == 0 || playPauseStatus == 2)
	{ 
		$("#play").attr("class","glyphicon glyphicon-play aligned")
		playPauseStatus = 1;
		isNotPaused = false;
	} 
	else if(playPauseStatus == 1) 
	{    
		$("#play").attr("class","glyphicon glyphicon-pause aligned")
		playPauseStatus = 2;
		isNotPaused = true;
		if(animationId == 2){
			staticBlockChainImageHandler();
		}
	}
}


function toggleImage(){
	animationId++;
	if(animationId == 1 ){
		var frame1 = 1;
		anim1 = setInterval(
			function(){
				if(isNotPaused){
					$('#displayAnims').attr('src',"../app/app-content/images/Frames/A1_Frame"+frame1+".png");
					frame1++;
					if(frame1 == 2){
						$('.animationGif').append("<div><img src='../app/app-content/images/Callouts/11.png' class='callout1' /></div>")
					}
					if(frame1 == 14){
						$('.animationGif').append("<div><img src='../app/app-content/images/Callouts/12.png' class='callout2'  /></div>")
					}
					if(frame1 == 19){
						$('.animationGif').append("<div><img src='../app/app-content/images/Callouts/13.png' class='callout3' /></div>")
					}
					if(frame1==22){
						clearInterval(anim1);
						$( "div img" ).remove( ".callout1");
						$( "div img" ).remove( ".callout2");
						$( "div img" ).remove( ".callout3");
						toggleImage();		
					}
				}
			},600)
	}
	else if(animationId==2){
		retrieveBlockChainId()
		staticBlockChainImageHandler();
	}
	else if(animationId==3){
		$('#displayAnims').attr('src',"../app/app-content/images/Frames2/A2_Fram1.png");
		$('.staticImage').fadeOut(500, function() {
			var frame2 = 1;
			anim2 = setInterval(
				function(){
					if(isNotPaused){
						$('#displayAnims').attr('src',"Frames2/A2_Fram"+frame2+".png");
						frame2++;
						if(frame2 == 10){
							$('.animationGif').append("<div><img src='../app/app-content/images/Callouts/15.png' class='callout5' /></div>")
						}
						if(frame2==12){
							clearInterval(anim2);
							redirectToApplication();
						}
					}
				},600)
			$('.animationGif').fadeIn(500);
		});
	}
}

function staticBlockChainImageHandler(){
	$('.animationGif').fadeOut(500, function() {
		$('.staticImage').fadeIn(500,function(){
			$(".blockArrow").delay(2500).queue(function() {
				if(isNotPaused){
					$(".blockArrow").fadeIn(500,function(){
						if(isNotPaused){
							$('.staticImage').append("<div><img src='../app/app-content/images/Callouts/14.png' class='callout4' /></div>")
							$(this).delay(500).queue(function() {
								if(isNotPaused){
									$( "div img" ).remove( ".callout4");
									toggleImage();
								}else{
									$(".blockArrow").dequeue();
								}	
							});
						}
					});
				}
				$(".blockArrow").dequeue();
			});	
		});
	});
}

function redirectToApplication(){
	//write code here to redirect to main application
	$( "div img" ).remove( ".callout5");
}


//dummy function as per  test JSON
function retrieveTransactionHash() {
    var url = "http://localhost:3000/transaction"
    $.ajax({
        type: "GET",
        url: url,
        dataType: "JSON",
        success: function(data){
           $("#blockHashCodeValue").text(data.id);
        },
        error: function(){
        }
    });
}

//dummy function as per  test JSON
function retrieveBlockChainId() {
    var url = "http://localhost:3000/blockCode"
    $.ajax({
        type: "GET",
        url: url,
        dataType: "JSON",
        success: function(data){
           $(".blockTextPlacement1").text(data.number);
		   $(".blockTextPlacement2").text(data.number - 1 );
		   $(".blockTextPlacement3").text(data.number - 2 );
		   $(".blockTextPlacement4").text(data.number - 3 );
		   $(".blockText").show();
        },
        error: function(){
        }
    });
}