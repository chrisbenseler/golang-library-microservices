import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BooksFormComponent } from './pages/books-form/books-form.component';
import { BooksComponent } from './pages/books/books.component';
import { ReviewsFormComponent } from './pages/reviews-form/reviews-form.component';
import { ReviewsComponent } from './pages/reviews/reviews.component';

const routes: Routes = [
  { path: 'books', component: BooksComponent },
  { path: 'books/:id/reviews', component: ReviewsComponent },
  { path: 'books/:id/reviews/new', component: ReviewsFormComponent },
  { path: 'books/new', component: BooksFormComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
