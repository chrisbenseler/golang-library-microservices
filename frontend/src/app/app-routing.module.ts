import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BooksFormComponent } from './pages/books-form/books-form.component';
import { BooksComponent } from './pages/books/books.component';

const routes: Routes = [
  { path: 'books', component: BooksComponent },
  { path: 'books/new', component: BooksFormComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
