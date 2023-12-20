var invoiceList =[]
let shop=''
let owner=''


document.addEventListener('DOMContentLoaded', function () {
    populateYears();
    populateMonths();
    document.getElementById("showHistory").addEventListener("click", function () {
        const month = document.getElementById("monthSelect").value;
        const year = document.getElementById("yearSelect").value;
        const currentUrl = window.location.href;
        const pathParts = currentUrl.split('/');
        shop = decodeURIComponent(pathParts[pathParts.indexOf('details') + 1]);
        owner = decodeURIComponent(pathParts[pathParts.indexOf('details') + 2]);

        // Log or use the extracted values

        const formData = {
            Month: month,
            Year: year,
        };

        console.log("Shop:", shop);
        console.log("Owner:", owner);
        const url = `/customer/details/${shop}/${owner}`;

        fetch(url, {
            method: 'POST',
            body: JSON.stringify(formData),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);
                if (data.GetSuccess) {
                    invoiceList=data.InvoiceList
                    showPurchasedTable()
                } else {
                    Swal.fire({
                        icon: 'error',
                        title: 'Record not Found',
                        text: 'There was an error adding the customer. Please try again.',
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
});

function populateYears() {
    // Get the current year
    const currentYear = new Date().getFullYear();

    // Select the <select> element for years
    const yearSelect = document.getElementById('yearSelect');

    // Populate the select element with the last three years
    for (let i = 0; i < 3; i++) {
        const option = document.createElement('option');
        const year = currentYear - i;
        option.value = year;
        option.textContent = year;
        yearSelect.appendChild(option);
    }
}

// Function to populate months in the select element
function populateMonths() {
    // Array of month names
    const months = [
        'January', 'February', 'March', 'April', 'May', 'June',
        'July', 'August', 'September', 'October', 'November', 'December'
    ];

    // Select the <select> element for months
    const monthSelect = document.getElementById('monthSelect');

    // Populate the select element with month names
    months.forEach(month => {
        const option = document.createElement('option');
        option.value = month; // You can set a specific value if needed
        option.textContent = month;
        monthSelect.appendChild(option);
    });
}

function showPurchasedTable() {
    const tableBody = document.getElementById("purchasedTable").getElementsByTagName('tbody')[0];
    tableBody.innerHTML = "";

    for (let i = 0; i < invoiceList.length; i++) {
        const invoice = invoiceList[i];

        const row = tableBody.insertRow(tableBody.rows.length);
        const checkboxCell = row.insertCell(0);
        const date = row.insertCell(1);
        const invoiceNo = row.insertCell(2);
        const grandTotalBill = row.insertCell(3);
        const action = row.insertCell(4);

        checkboxCell.innerHTML = '<label class="checkboxs"><input type="checkbox"><span class="checkmarks"></span></label>';
        date.innerText = invoice.Date;
        invoiceNo.innerText = invoice.InvoiceNo;
        grandTotalBill.innerText = invoice.GrandTotal;

        const actionDiv = document.createElement('div');
        actionDiv.classList.add('d-flex');

        const viewCustomerLink = document.createElement('a');
        viewCustomerLink.href = 'javascript:void(0);';
        viewCustomerLink.classList.add('view-invoice');
        viewCustomerLink.dataset.rowNumber = i; // Store the row number as a data attribute
        viewCustomerLink.innerHTML = '<img src="../../../assets/img/icons/eye.svg" alt="svg">';
        actionDiv.appendChild(viewCustomerLink);

        const spacerDiv = document.createElement('div');
        spacerDiv.classList.add('mx-2');
        actionDiv.appendChild(spacerDiv);

        const deleteCustomerLink = document.createElement('a');
        deleteCustomerLink.href = 'javascript:void(0);';
        deleteCustomerLink.classList.add('delete-invoice');
        deleteCustomerLink.innerHTML = '<img src="../../../assets/img/icons/delete.svg" alt="svg">';
        actionDiv.appendChild(deleteCustomerLink);

        // Append the customized action content to the action cell
        action.appendChild(actionDiv);
    }

    // Attach event listener to all elements with class 'view-invoice'

    const viewInvoiceLinks = document.querySelectorAll('.view-invoice');
    viewInvoiceLinks.forEach(link => {
        link.addEventListener('click', function () {
            const rowNumber = this.dataset.rowNumber;
            console.log(`View invoice clicked for row number ${rowNumber}`);

            fetchInvoiceData('../../../sales/invoice',invoiceList[rowNumber])
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.text();
                })
                .then(htmlContent => {
                    console.log('Successfully fetched index.html');
                    const newWindow = window.open();
                    newWindow.document.write(htmlContent);
                    newWindow.document.close();
                })
                .catch(error => console.error('Error fetching HTML file:', error));
        });
    });

}

function fetchInvoiceData(url,invoiceData) {
    return fetch(url, {
        method: 'POST',
        body: JSON.stringify(invoiceData),
        headers: {
            'Content-Type': 'application/json'
        }
    });
}