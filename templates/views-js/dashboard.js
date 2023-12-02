document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("click_logout").addEventListener("click", function () {

        fetch('/signout', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/signin';
                } else {
                    console.error('Logout failed');
                }
            })
            .catch(error => console.error('Error during logout:', error));
    });
});
