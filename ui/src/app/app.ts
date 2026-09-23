import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Monitors } from './pages/monitors/monitors';

@Component({
  selector: 'app-root',
  imports: [
    RouterOutlet,
    Monitors,
  ],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})
export class App {
  protected readonly title = signal('app');
}
