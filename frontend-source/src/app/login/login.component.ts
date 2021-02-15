import { Component, OnInit } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {MatIcon} from "@angular/material/icon";
import {MatFormField} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatSelect} from "@angular/material/select";
import {MatOption} from "@angular/material/core";
import {FormControl} from "@angular/forms";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {Router} from "@angular/router";
import {getLoginURL} from "../app.component";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
@Injectable()
export class LoginComponent implements OnInit {
  loginForm : FormGroup;
  selected = 'fe3hau';
  constructor(private http: HttpClient, private formBuilder: FormBuilder, private router: Router) {
    this.loginForm = this.formBuilder.group({
      name: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20), Validators.pattern("^[a-zA-Z0-9]*$")]],
      game: ['', Validators.required]
    });
  }

  join(user: string, game: string) : void {
    user = user.toLowerCase();
    const url : string = getLoginURL() + "/" + game + "/user/" + user + "/create";
    this.http.get(url).subscribe((data : {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if(data["status"] == 0) {
        console.log("logging in");
        this.router.navigate(['./game', game, user]);
      }
    })
    console.log(user + " has logged into " + game);
  }

  ngOnInit(): void {
  }

}
