import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.scss']
})
export class ReviewsComponent implements OnInit {

  bookId: string;

  reviews$: any;

  constructor(private route: ActivatedRoute, private apiService: ApiService) {
    this.route.params.subscribe(params => {
      this.bookId = params.id;
    })
  }

  ngOnInit() {
    this.reviews$ = this.apiService.reviews(this.bookId);
  }

}
