document.getElementById("click_sign_up").addEventListener("click", function () {
    const fullName = document.querySelector('input[name="full_name"]').value;
    const email = document.querySelector('input[name="email_add"]').value;
    const password = document.querySelector('input[name="password"]').value;

    // Create a data object to hold the form data
    const formData = {
        full_name: fullName,
        email_add: email,
        password: password
    };

    // Send the data to the Go backend using a POST request
    fetch('/signup', {
        method: 'POST',
        body: JSON.stringify(formData),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log("Response from server:", data);

            if (!data.full_name.valid) {
                console.log("Full Name Invalid")
                document.querySelector('input[name="full_name"]').classList.add('is-invalid');
            }
            if (!data.email_add.valid) {
                console.log("Email Invalid")
                document.querySelector('input[name="email_add"]').classList.add('is-invalid');
            }
            if (!data.password.valid) {
                console.log("Password Invalid")
                document.querySelector('input[name="password"]').classList.add('is-invalid');
            }
            if (data.full_name.valid && data.email_add.valid && data.password.valid) {
                // Show the success modal
                console.log("Show Modal")
                $('#successModal').modal('show');
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});

$(document).ready(function() {
    document.getElementById("click_modal_close").addEventListener("click", function() {
        const modal = document.getElementById("successModal");
        $(modal).modal("hide");
        window.location.href = "/signup";

    });
});

