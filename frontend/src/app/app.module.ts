import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BooksComponent } from './pages/books/books.component';
import { HttpClientModule } from '@angular/common/http';
import { BooksFormComponent } from './pages/books-form/books-form.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ReviewsComponent } from './pages/reviews/reviews.component';
import { ReviewsFormComponent } from './pages/reviews-form/reviews-form.component';

@NgModule({
  declarations: [
    AppComponent,
    BooksComponent,
    BooksFormComponent,
    ReviewsComponent,
    ReviewsFormComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
