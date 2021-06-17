import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {
  title = 'client';
  transactions: Transaction[] = [];
  transactionForm = this.formBuilder.group({
    id: "",
    type: "",
    desc: "",
    amount: 0,
    date: new Date(),
  })

  constructor(private formBuilder: FormBuilder) {}

  ngOnInit(): void {
    this.getTransactions()
  }

  getTransactions() {
    fetch("http://localhost:8080/transaction", {
      method: "GET",
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
    }).then(res => res.json().then(transactions => this.transactions = transactions))
  }

  addTransaction() {
    let newTransaction: Transaction = this.transactionForm.value as Transaction

    fetch("http://localhost:8080/transaction", {
      method: "POST",
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(newTransaction)
    }).then(() => this.transactions.push(newTransaction))
  }

  updateTransaction() {
    let updatedTransaction: Transaction = this.transactionForm.value as Transaction

    fetch("http://localhost:8080/transaction", {
      method: "PUT",
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(updatedTransaction)
    }).then(() => this.transactions.push(updatedTransaction))
  }
}
