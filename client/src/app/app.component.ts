import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {
  title = "Angular Go CRUD";
  error: string | null = null;

  transactions: Transaction[] = [];
  transactionForm = this.formBuilder.group({
    id: null,
    type: "",
    desc: "",
    amount: 0,
    date: new Date(),
  })

  constructor(private formBuilder: FormBuilder) {
  }

  ngOnInit(): void {
    this.getTransactions()
  }

  getTransactions() {
    fetch("http://localhost:8080/transaction", {
      method: "GET",
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
    }).then(res => {
      res.json().then(data => {
        if (!res.ok) {
          this.error = "LOAD: " + data;
          return;
        }

        this.transactions = data
        this.error = null;
      })
    })
  }

  addTransaction() {
    let newTransaction: Transaction = this.transactionForm.value as Transaction
    newTransaction.date = new Date(newTransaction.date).toISOString(),

      fetch("http://localhost:8080/transaction", {
        method: "POST",
        headers: {'Content-Type': 'application/json;charset=UTF-8'},
        body: JSON.stringify(newTransaction)
      }).then(res => {
        res.json().then(data => {
          if (!res.ok) {
            this.error = "ADD: " + data;
            return;
          }

          newTransaction.id = data["id"]
          this.transactions.push(newTransaction)
          this.error = null;
        })
      })
  }

  updateTransaction() {
    let updatedTransaction: Transaction = this.transactionForm.value as Transaction
    updatedTransaction.date = new Date(updatedTransaction.date).toISOString(),

      fetch("http://localhost:8080/transaction", {
        method: "PUT",
        headers: {'Content-Type': 'application/json;charset=UTF-8'},
        body: JSON.stringify(updatedTransaction)
      }).then(res => {
        if (!res.ok) {
          res.text().then(text => this.error = "UPDATE: " + text);
          return;
        }

        this.transactions.push(updatedTransaction);
        this.error = null;
      })
  }

  deleteTransaction() {
    let transactionId = this.transactionForm.value.id

    fetch("http://localhost:8080/transaction/" + transactionId, {
      method: "DELETE",
      headers: {'Content-Type': 'application/json;charset=UTF-8'}
    }).then(res => {
      if (!res.ok) {
        res.text().then(text => this.error = "DELETE: " + text);
        return;
      }

      this.transactions = this.transactions.filter(transaction => transaction.id != transactionId)
      this.error = null;
    })
  }
}
