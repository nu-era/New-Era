/**
 *********************************
TODO:

    1. Send post request to users
        a. On success, store auth key and possibly user struct in local storage, redirect to lobbies
        b. On fail, show alert with error message
    2. Send get request to users (or whatever validation is)
        a. On success, store auth key and possibly user struct in local storage, redirect to lobbies
        b. On fail, show alert with error message

 *********************************
 */

 window.onload = function() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(function (position) {
            localStorage.setItem('lat', position.coords.latitude.toString())
            localStorage.setItem('long', position.coords.longitude.toString())
        }, function () {
            console.log("error getting location")
        });
    } else {
        var lat = "47.655548"
        var long = "-122.303200"
        localStorage.setItem('lat', lat)
        localStorage.setItem('long', long)
        console.log('no geolocation support')
    }
};



const newUserUrl = "https://api.bfranzen.me/setup"
const retUserUrl = "https://api.bfranzen.me/connect"


$('.form').find('input, textarea').on('keyup blur focus', function (e) {

    var $this = $(this),
        label = $this.prev('label');

    if (e.type === 'keyup') {
        if ($this.val() === '') {
            label.removeClass('active highlight');
        } else {
            label.addClass('active highlight');
        }
    } else if (e.type === 'blur') {
        if ($this.val() === '') {
            label.removeClass('active highlight');
        } else {
            label.removeClass('highlight');
        }
    } else if (e.type === 'focus') {

        if ($this.val() === '') {
            label.removeClass('highlight');
        }
        else if ($this.val() !== '') {
            label.addClass('highlight');
        }
    }

});

$('.tab a').on('click', function (e) {
    e.preventDefault();

    $(this).parent().addClass('active');
    $(this).parent().siblings().removeClass('active');

    target = $(this).attr('href');

    $('.tab-content > div').not(target).hide();

    $(target).fadeIn(600);

});

// Logic to send new user or returning user data to server

$('#new-user-form').submit(function (e) {
    e.preventDefault();
    var formInputs = $('#new-user-form :input');

    var values = {};
    formInputs.each(function () {
        values[this.name] = $(this).val();
    });
    values.latitude = localStorage.getItem('lat');
    values.longitude = localStorage.getItem('long');
    
    var valJson = JSON.stringify(values);

    $.ajax({
        method: "POST",
        url: newUserUrl,
        contentType: 'application/json',
        data: valJson,
        success: function (data, textStatus, response) {
            var auth = response.getResponseHeader('Authorization');
            var pubKey = response.getResponseHeader('X-VapidKey');
            var userData = JSON.stringify(data);
            localStorage.setItem('phone', values.phone);
            localStorage.setItem('auth', auth);
            localStorage.setItem('device', userData);
            localStorage.setItem('pubKey', pubKey);
            window.location.replace("./html/verification.html");
        },
        error: function (jqXhr, textStatus, errorThrown) {
           // alert(jqXhr.responseText);
        }
    })

});

$('#user-form').submit(function (e) {
    e.preventDefault();
    var formInputs = $('#user-form :input');
    var values = {};
    formInputs.each(function () {
        values[this.name] = $(this).val();
    });
    values.latitude = localStorage.getItem('lat');
    values.longitude = localStorage.getItem('long');
    console.log(values);

    var valJson = JSON.stringify(values);
    $.ajax({
        method: "POST",
        url: retUserUrl,
        contentType: 'application/json',
        data: valJson,
        success: function (data, textStatus, response) {
            var auth = response.getResponseHeader('Authorization');
            var pubKey = response.getResponseHeader('X-VapidKey');
            var userData = JSON.stringify(data);
            localStorage.setItem('auth', auth);
            localStorage.setItem('device', userData);
            localStorage.setItem('pubKey', pubKey);
            //switchToVerification();
            window.location.replace("./html/alert.html");
        },
        error: function (jqXhr, textStatus, errorThrown) {
           // alert(jqXhr.responseText);
        }
    })
});

function switchToVerification() {

}
