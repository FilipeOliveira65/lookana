import { Component, Input } from '@angular/core';
import { Output, EventEmitter } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { ApiService } from '../../core/services/api.service';
import { response } from 'express';
import { error } from 'console';

@Component({
  selector: 'app-new-service-window',
  standalone: true,
  imports: [
    ReactiveFormsModule
  ],
  templateUrl: './new-service-window.html',
  styleUrl: './new-service-window.scss',
})
export class NewServiceWindow {
  @Output() closeWindow = new EventEmitter<boolean>();
  @Input() visible = false;

  constructor(
    private apiService: ApiService
  ){}

  serviceForm = new FormGroup ({
    serviceName: new FormControl(''),
    serviceDns: new FormControl('')
  })

  closeNewServiceWindow(): void {
    this.closeWindow.emit(false);
  }

  createNewService(): void {
    var serviceNameValue = this.serviceForm.value!.serviceName!;
    var serviceDnsValue = this.serviceForm.value!.serviceDns!;

    var newServiceJson = {
      "service-name": serviceNameValue,
      "service-dns": serviceDnsValue,
    }

    this.apiService.newServiceDataToTheApi(newServiceJson)
      .subscribe({
        next: (response) => {
          console.log('Service created. ', response)
        },
        error: (error) => {
          console.log("Wasn't possible to connect with the api", error)
        }
      })

    console.log(`Service name: ${serviceNameValue}; Service DNS: ${serviceDnsValue}`);

  }
}
