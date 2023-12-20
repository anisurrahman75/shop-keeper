document.addEventListener('DOMContentLoaded', function () {
    document.getElementById("submit").addEventListener("click", function () {
        const brandName= document.getElementById("brand").value
        const unitName= document.getElementById("unit").value

        const formData = {
            Brand :  brandName,
            Name : document.getElementById("productName").value,
            Grade: document.getElementById("productGrade").value,
            Unit: unitName,
            Description: document.getElementById("description").value,
        };

        console.log(formData)

        // Send the data to the Go backend using a POST request
        fetch('/product/add', {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);

                if (data.AddSuccess) {
                    // Brand added successfully, show success alert
                    Swal.fire({
                        icon: 'success',
                        title: 'Product Added Successfully',
                        text: 'The Product has been added to the database.',
                    }).then(() => {
                        // Reset specific form fields
                        // resetFormFields();
                        // Redirect to the specified URL
                        window.location.href = '/product/add';
                    });
                } else {
                    Swal.fire({
                        icon: 'error',
                        title: 'Product Not Added',
                        text: 'There was an error adding the product. Please try again.',
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    })

    function resetFormFields() {
        // Reset specific form fields
        const unitSelect = document.getElementById("unit");
        unitSelect.selectedIndex = 0;

        document.getElementById("brand").selectedIndex = 0;
        // document.getElementById("brand").selectedIndex = 1;
        // document.getElementById("unit").selectedIndex = 1;
        document.getElementById("productName").value = '';
        document.getElementById("productGrade").value = '';
        document.getElementById("description").value = '';
    }
});
