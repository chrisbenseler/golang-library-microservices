import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.scss']
})
export class SignInComponent implements OnInit {

  form: FormGroup;

  errorMessage: string = null;
  loading = false;

  constructor(private fb: FormBuilder, private apiService: ApiService, private router: Router) {
    this.form = this.fb.group({
      email: ["", [Validators.required, Validators.email]],
      password: ["", [Validators.required]],
    });
  }

  ngOnInit() {}

  onSubmit() {
    this.errorMessage = null;
    this.loading = true;
    const { email, password } = this.form.value;
    this.apiService
      .signIn(email, password)
      .subscribe((result) => {
        this.loading = false;
        this.router.navigateByUrl('/books')
      }, (err: HttpErrorResponse) => {
        console.log(err)
        this.loading = false;
        this.errorMessage = err.error.error;
      });
  }

}
