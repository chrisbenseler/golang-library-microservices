import { Component, OnInit } from "@angular/core";
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { Router } from '@angular/router';
import { ApiService } from "src/app/services/api.service";

@Component({
  selector: "app-books-form",
  templateUrl: "./books-form.component.html",
  styleUrls: ["./books-form.component.scss"],
})
export class BooksFormComponent implements OnInit {
  form: FormGroup;

  constructor(private fb: FormBuilder, private apiService: ApiService, private router: Router) {
    this.form = this.fb.group({
      title: ["", [Validators.required]],
      year: ["", [Validators.required]],
    });
  }

  ngOnInit() {}

  onSubmit() {
    const values = this.form.value;
    this.apiService
      .createBook(values.title, parseInt(values.year))
      .subscribe((result) => {
        this.router.navigateByUrl('/books')
      });
  }
}
