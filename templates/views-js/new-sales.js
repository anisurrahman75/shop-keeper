let productsSet = new Set();
let productsList = [];
let totalBill = 0, discountByPercent =  0, afterDiscountTotal = 0, disCountSave = 0;
let customerInfo ={
    ShopNameOwner: '',
    Shop          : '',
    Owner         : '',
    PhoneNumber   : '',
    Address       : '',
    TotalDue      : 0,
}

let customerList = Array.of(customerInfo)

document.addEventListener('DOMContentLoaded', function () {
        // Send the data to the Go backend using a POST request
        fetch('/get/customer/list', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                customerList=data
                console.log("Response from server:", customerList);
                const selectElement = document.getElementById("ShopNameOwner");
                // Loop through the array and create options dynamically
                for (let i = 0; i < customerList.length; i++) {
                    const option = document.createElement("option");
                    option.text = customerList[i].ShopNameOwner;
                    // You can also set the value attribute if needed
                    // option.value = shopNames[i];
                    selectElement.add(option);
                }

            })
            .catch(error => {
                console.error("Error:", error);
            });
});

document.getElementById("showCustomerInfo").addEventListener("click", function () {
    const shopName = document.getElementById("ShopNameOwner");
    const selectedIndex = shopName.selectedIndex;
    if (selectedIndex!==0){
                customerInfo=customerList[selectedIndex-1]
                document.getElementById("ownerName").innerText = customerInfo.Owner;
                document.getElementById("address").innerText = customerInfo.Address;
                document.getElementById("phoneNumber").innerText = customerInfo.PhoneNumber;
                document.getElementById("totalDue").innerText ="Tk "+ customerInfo.TotalDue;
    }
});

document.getElementById("submit-product").addEventListener("click", function () {
    const productName = document.getElementById("productName").value;
    const unitPrice = document.getElementById("unitPrice").value;
    const qty = document.getElementById("qty").value;

    productsSet.add({ name: productName, unitPrice: parseFloat(unitPrice).toFixed(2), qty: qty, subtotal: (unitPrice * qty).toFixed(2) });
    productsList = Array.from(productsSet);

    updateTable();
    resetProductForm();
});

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
        CustomerInfo: customerInfo,
        Date: document.getElementById("date").value,
        InvoiceNo: parseInt(document.getElementById("invoiceNo").value).toString(),
        ProductsInfo: productsList,
        NetTotal:  totalBill.toFixed(2).toString(),
        DiscountInPercent: discountByPercent.toString(),
        SaveInDiscount: disCountSave.toFixed(2).toString(),
        GrandTotal: afterDiscountTotal.toFixed(2).toString()
    };

    console.log("Net Total: ", invoiceData.NetTotal);
    console.log("DisCount: ", invoiceData.DiscountInPercent);
    console.log("Save Discount: ", invoiceData.SaveInDiscount);
    console.log("Grand Total: ", invoiceData.GrandTotal);

    console.log('Fetching index.html...');

    fetchInvoiceData('/sales/invoice',invoiceData)
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

document.getElementById("addInvoice").addEventListener("click", function () {
    const invoiceData = {
        CustomerInfo: customerInfo,
        Date: document.getElementById("date").value,
        InvoiceNo: parseInt(document.getElementById("invoiceNo").value).toString(),
        ProductsInfo: productsList,
        NetTotal:  totalBill.toFixed(2).toString(),
        DiscountInPercent: discountByPercent.toString(),
        SaveInDiscount: disCountSave.toFixed(2).toString(),
        GrandTotal: afterDiscountTotal.toFixed(2).toString()
    };

    console.log("Net Total: ", invoiceData.NetTotal);
    console.log("DisCount: ", invoiceData.DiscountInPercent);
    console.log("Save Discount: ", invoiceData.SaveInDiscount);
    console.log("Grand Total: ", invoiceData.GrandTotal);

    fetchInvoiceData('/sales/invoice-add',invoiceData)
        .then(response => response.json())
        .then(data => {
            console.log("Response from server:", data);
            if (data.AddSuccess) {
                Swal.fire({
                    icon: 'success',
                    title: 'Customer Added Successfully',
                    text: 'The Customer has been added to the database.',
                }).then(() => {
                    window.location.href = '/sales/new';
                });
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'Customer Not Added',
                    text: 'There was an error adding the customer. Please try again.',
                });
            }
        }).catch(error => {
        console.error("Error:", error);
    });

});



function fetchInvoiceData(url,invoiceData) {
    return fetch(url, {
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
        productNameCell.innerHTML = `<a href="javascript:void(0);">${product.name}</a>`;
        priceCell.innerText = product.unitPrice;
        quantityCell.innerText = product.qty;
        totalCell.innerText = product.subtotal;

        totalBill += parseFloat(product.subtotal);
    }

    discountByPercent = parseFloat(document.getElementById("discount").value) || 0;
    afterDiscountTotal = totalBill - ((discountByPercent/100) * totalBill);
    disCountSave = totalBill - afterDiscountTotal;

    document.getElementById("netTotal").innerText = `Tk ${totalBill.toFixed(2)}`;
    document.getElementById("saveMoney").innerText = `Tk ${disCountSave.toFixed(2)}`;
    document.getElementById("grandTotal").innerText = `Tk ${afterDiscountTotal.toFixed(2)}`;
}


function resetProductForm() {
    document.getElementById("qty").value = "";
    document.getElementById("unitPrice").value = "";
    document.getElementById("productName").selectedIndex = 0;
}