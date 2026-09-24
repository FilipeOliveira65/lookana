import { Component, Input } from '@angular/core';
import { Monitors } from '../../pages/monitors/monitors';

@Component({
  selector: 'app-new-service-window',
  standalone: true,
  imports: [
    Monitors
  ],
  templateUrl: './new-service-window.html',
  styleUrl: './new-service-window.scss',
})
export class NewServiceWindow {
  @Input() visible = false;

   constructor(
    private monitors: Monitors
  ){}

  closeNewServiceWindow(): void {
    this.monitors.newServiceWindowIsVisible = false
  }
}
