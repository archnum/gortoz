/*
####### gortoz (c) 2025 Archivage Numérique ######################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

document.addEventListener('DOMContentLoaded', () => {
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

    $navbarBurgers.forEach( el => {
        el.addEventListener('click', () => {
            const target = el.dataset.target;
            const $target = document.getElementById(target);
            el.classList.toggle('is-active');
            $target.classList.toggle('is-active');
        });
    });
});

document.body.addEventListener('htmx:beforeSwap', function(evt) {
    if (evt.detail.xhr.status === 200 || evt.detail.xhr.status === 204) {
        document.getElementById("error").style.display = "none";
    } else {
        document.getElementById("error").style.display = "block";
        evt.detail.shouldSwap = true;
        evt.detail.target = htmx.find("#error");
    }
});

/*
####### END ############################################################################################################
*/