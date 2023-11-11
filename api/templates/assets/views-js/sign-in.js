document.getElementById("click_sign_in").addEventListener("click", function () {
    const email = document.querySelector('input[name="email_add"]').value;
    const password = document.querySelector('input[name="password"]').value;

    // Create a data object to hold the form data
    const formData = {
        email_add: email,
        password: password
    };

    console.log("email:",email)
    console.log("pass:",password)

    // Send the data to the Go backend using a POST request
    fetch('/signin', {
        method: 'POST',
        body: JSON.stringify(formData),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log("Response from server:", data);
            if (data.authorize_user) {
                console.log("Login Successfully")
                window.location.href = "/dashboard";
            }
            else {
                console.log("Incorrect UserName or Password")
            }

        })
        .catch(error => {
            console.error("Error:", error);
        });
});
