document.addEventListener('DOMContentLoaded', function () {
    let productsSet = new Set();
    let productsList = [];
    let totalBill = 0;

    function fetchInvoiceData(invoiceData) {
        return fetch('/invoice', {
            method: 'POST',
            body: JSON.stringify(invoiceData),
            headers: {
                'Content-Type': 'application/json'
            }
        });
    }

    function updateTable() {
        const tableBody = document.getElementById("productTable").getElementsByTagName('tbody')[0];
        tableBody.innerHTML = "";

        totalBill = 0;

        for (const product of productsList) {
            const row = tableBody.insertRow(tableBody.rows.length);

            const deleteCell = row.insertCell(0);
            const productNameCell = row.insertCell(1);
            const priceCell = row.insertCell(2);
            const quantityCell = row.insertCell(3);
            const totalCell = row.insertCell(4);

            deleteCell.innerHTML = '<a class="delete-set"><img src="../assets/img/icons/delete.svg" alt="svg"></a>';
            productNameCell.innerHTML = '<a href="javascript:void(0);">' + product.name + '</a>';
            priceCell.innerText = product.unitPrice;
            quantityCell.innerText = product.qty;
            totalCell.innerText = product.subtotal;

            totalBill += product.subtotal;
        }

        document.getElementById("grandTotal").innerText = "$ " + totalBill;
    }

    document.getElementById("showCustomerInfo").addEventListener("click", function () {
        const shopName = document.getElementById("shopName").value;
        fetch(`/customer/${shopName}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                console.log("Response from server:", data);
                document.getElementById("ownerName").innerText = data.owner_name;
                document.getElementById("address").innerText = data.address;
                document.getElementById("phoneNumber").innerText = data.phone_number;
                document.getElementById("totalDue").innerText = data.total_due;
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });

    document.getElementById("submit-product").addEventListener("click", function () {
        const productName = document.getElementById("productName").value;
        const unitPrice = document.getElementById("unitPrice").value;
        const qty = document.getElementById("qty").value;

        productsSet.add({ name: productName, unitPrice: unitPrice, qty: qty, subtotal: unitPrice * qty });
        productsList = Array.from(productsSet);

        addRowsToTable();
        resetProductForm();
    });

    function addRowsToTable() {
        updateTable();
    }

    function resetProductForm() {
        document.getElementById("qty").value = "";
        document.getElementById("unitPrice").value = "";
        document.getElementById("productName").selectedIndex = 0;
    }

    document.getElementById("productTable").addEventListener("click", function (event) {
        const target = event.target;

        if (target.tagName === "IMG" && target.getAttribute("src") === "../assets/img/icons/delete.svg") {
            const rowIndex = target.closest("tr").rowIndex;
            const deletedProduct = productsList.splice(rowIndex - 1, 1)[0];
            productsSet = new Set(productsList);
            updateTable();
        }
    });

    document.getElementById("printInvoice").addEventListener("click", function () {
        const invoiceData = {
            ShopName: document.getElementById("shopName").value,
            OwnerName: document.getElementById("ownerName").innerText,
            ProductsInfo: productsList,
        };

        console.log("Owner Name: ", invoiceData.OwnerName);

        console.log('Fetching index.html...');

        fetchInvoiceData(invoiceData)
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
