import {Component} from '@angular/core';
import {variable} from "@angular/compiler/src/output/output_ast";

interface Transaction {
  type: string;
  desc: string;
  amount: number;
  date: Date
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  title = 'client';

  async sendMessage(msg: string) {
    let body: Transaction = {
      "type": "Income",
      "desc": "Investment Interest",
      "amount": 75.00,
      "date": new Date(),
    }

    const response = await fetch("http://localhost:8080/transaction", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json;charset=UTF-8'
      },
      body: JSON.stringify(body)
    })
  }
}
