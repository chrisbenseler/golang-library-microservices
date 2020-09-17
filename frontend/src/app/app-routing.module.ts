import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BooksFormComponent } from './pages/books-form/books-form.component';
import { BooksComponent } from './pages/books/books.component';
import { ReviewsFormComponent } from './pages/reviews-form/reviews-form.component';
import { ReviewsComponent } from './pages/reviews/reviews.component';
import { SignInComponent } from './pages/sign-in/sign-in.component';

const routes: Routes = [
  { path: 'books', component: BooksComponent },
  { path: 'books/:id/reviews', component: ReviewsComponent },
  { path: 'books/:id/reviews/new', component: ReviewsFormComponent },
  { path: 'books/new', component: BooksFormComponent },
  { path: 'auth/signin', component: SignInComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
