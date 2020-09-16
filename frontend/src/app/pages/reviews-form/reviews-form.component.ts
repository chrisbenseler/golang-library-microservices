import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-reviews-form',
  templateUrl: './reviews-form.component.html',
  styleUrls: ['./reviews-form.component.scss']
})
export class ReviewsFormComponent implements OnInit {

  form: FormGroup;

  bookId: string;

  constructor(private fb: FormBuilder, private apiService: ApiService, private router: Router, private route: ActivatedRoute) {
    this.form = this.fb.group({
      content: ["", [Validators.required]]
    });

    this.route.params.subscribe(params => {
      this.bookId = params.id;
    })

  }

  ngOnInit() {}

  onSubmit() {
    const values = this.form.value;
    this.apiService
      .createReview("books", this.bookId, values.content)
      .subscribe((result) => {
        this.router.navigateByUrl('/books/' + this.bookId + '/reviews')
      });
  }

}
