document.addEventListener('DOMContentLoaded', function () {
    const apiUrl = '/product/brand-unit';
    document.getElementById("click-add-brand").addEventListener("click", function () {
        const formData = {
            Type: "brand",
            Name: document.getElementById("brand-name").value,
        };

        // Send the data to the Go backend using a POST request
        fetch(apiUrl, {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);

                if (data.AddSuccessful) {
                    // Brand added successfully, show success alert
                    Swal.fire({
                        icon: 'success',
                        title: 'Brand Added Successfully',
                        text: 'The brand has been added to the database.',
                    }).then(() => {
                        // Reload the page after the user clicks "OK"
                        // document.getElementById("brand-name").,
                        // location.reload();

                        window.location.href = apiUrl;
                    });
                } else {
                    // Brand not added, show error alert
                    Swal.fire({
                        icon: 'error',
                        title: 'Brand Not Added',
                        text: 'There was an error adding the brand. Please try again.',
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });


    document.getElementById("click-add-unit").addEventListener("click", function () {
        const formData = {
            Type: "unit",
            Name: document.getElementById("unit-name").value,
        };


        // Send the data to the Go backend using a POST request
        fetch(apiUrl, {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);

                if (data.AddSuccessful) {
                    // Brand added successfully, show success alert
                    Swal.fire({
                        icon: 'success',
                        title: 'Unit Added Successfully',
                        text: 'The Unit has been added to the database.',
                    }).then(() => {
                        // Reload the page after the user clicks "OK"
                        // document.getElementById("brand-name").,
                        // location.reload();

                        window.location.href = apiUrl;
                    });
                } else {
                    // Brand not added, show error alert
                    Swal.fire({
                        icon: 'error',
                        title: 'Unit Not Added',
                        text: 'There was an error adding the brand. Please try again.',
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });


    const deleteBrandIcons = document.querySelectorAll('.delete-brand');
    deleteBrandIcons.forEach(icon => {
        icon.addEventListener('click', function (event) {
            handleDelete('brand', event);
        });
    });

    const deleteUnitIcons = document.querySelectorAll('.delete-unit');
    deleteUnitIcons.forEach(icon => {
        icon.addEventListener('click', function (event) {
            handleDelete('unit', event);
        });
    });


    function handleDelete(type, event) {
        // Confirm with the user before proceeding with deletion
        Swal.fire({
            title: 'Are you sure?',
            text: 'You won\'t be able to revert this!',
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) {
                let formData={}
                switch (type) {
                    case "brand":
                        // If user confirms, proceed with deletion
                        const brandName = event.target.closest('tr').querySelector('td:nth-child(2)').textContent;
                        formData = {
                            Type: "brand",
                            Name: brandName,
                        };
                        break
                    case "unit":
                        const unitname = event.target.closest('tr').querySelector('td:nth-child(2)').textContent;
                        formData = {
                            Type: "unit",
                            Name: unitname,
                        };
                        break
                }

                console.log(formData)

                // Use fetch to send a DELETE request to your backend
                fetch(apiUrl, {
                    method: 'DELETE',
                    body: JSON.stringify(formData),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.DeleteSuccessful) {
                            // Show success alert
                            Swal.fire({
                                icon: 'success',
                                title: `${type.charAt(0).toUpperCase() + type.slice(1)} Deleted`,
                                text: `The ${type} has been deleted successfully.`,
                            }).then(() => {
                                // Reload the page after deletion
                                location.reload();
                            });
                        } else {
                            // Show error alert if deletion fails
                            Swal.fire({
                                icon: 'error',
                                title: 'Error',
                                text: `Failed to delete the ${type}. Please try again.`,
                            });
                        }
                    })
                    .catch(error => {
                        console.error('Error:', error);
                        // Show error alert if there's an issue with the fetch request
                        Swal.fire({
                            icon: 'error',
                            title: 'Error',
                            text: 'An unexpected error occurred. Please try again.',
                        });
                    });
            }
        });
    }


});
