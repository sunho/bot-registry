import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';
import { NgForm, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-form-input',
  templateUrl: './form-input.component.html',
  encapsulation: ViewEncapsulation.None,
  styles: ['app-form-input { display: block; }'],
})
export class FormInputComponent implements OnInit {
  @Input()
  type: string;
  @Input()
  name: string;
  @Input()
  form: NgForm;
  @Input()
  formGroup: FormGroup;

  constructor() {}

  ngOnInit(): void {}
}