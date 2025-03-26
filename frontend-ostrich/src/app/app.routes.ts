import { Routes } from '@angular/router';
import { FeedComponent } from './components/feed/feed/feed.component';
import { LoginComponent } from './components/login/login.component';

export const routes: Routes = [
    { path: '', redirectTo: 'feed', pathMatch: 'full' },
    { path: 'feed', component: FeedComponent },
    { path: 'login', component: LoginComponent }
];
