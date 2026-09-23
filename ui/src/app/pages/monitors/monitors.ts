import { Component } from '@angular/core';
import { NewServiceWindow } from '../../components/new-service-window/new-service-window';

@Component({
  selector: 'app-monitors',
  imports: [
    NewServiceWindow
  ],
  templateUrl: './monitors.html',
  styleUrl: './monitors.scss',
})
export class Monitors {}
