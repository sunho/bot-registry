import { Observable } from 'rxjs';
import { ModalService } from './../../services/modal.service';
import { PopupService } from './../../../core/services/popup.service';
import { BotService, CONFLICT, BAD_FORMAT } from './../../services/bot.service';
import { Component, OnInit } from '@angular/core';
import { Volume, Bot } from '../../models/bot';
import { ActivatedRoute } from '@angular/router';
import { STRINGS } from '../../../../locale/strings';

@Component({
  selector: 'app-volume',
  templateUrl: './volume.component.html',
})
export class VolumeComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private botService: BotService,
    private popupService: PopupService,
    private modalService: ModalService
  ) {}

  current: Bot;

  resNames = ['name', 'path'];

  resKeys = ['name', 'path'];

  resOptions = [{ title: 'Edit', func: this.editCallback.bind(this) }, { title: 'Delete', func: this.deleteCallback.bind(this) }];

  resItems: Volume[];

  deleteCallback(vol: Volume, s: string): void {
    this.botService.deleteVolume(this.current.id, vol.name).subscribe(
      _ => {
        this.refreshItems();
      },
      error => {
        this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
      }
    );
  }

  addCallback(obj: object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.addVolume(this.current.id, obj as Volume).subscribe(
        _ => {
          this.refreshItems();
          observer.next(true);
          observer.complete();
        },
        error => {
          if (error === CONFLICT) {
            this.popupService.createMsg(STRINGS.EXIST_VOLUME);
          } else if (error === BAD_FORMAT) {
            this.popupService.createMsg(STRINGS.BAD_VOLUME_FORMAT);
          } else {
            this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
          }
        }
      );
    });
  }

  editCallback(volume: Volume, s: string): void {
    this.modalService.createMod({
      title: 'Edit Volume',
      items: [
        {
          name: 'name',
          key: 'name',
          initial: volume.name,
          disabled: true
        },
        {
          name: 'path',
          key: 'path',
          initial: volume.path
        },
      ],
      button: 'Edit',
      callback: this.editCompleteCallback.bind(this),
    });
  }

  editCompleteCallback(obj: Object): Observable<boolean> {
    return new Observable<boolean>(observer => {
      this.botService.patchVolume(this.current.id, obj as Volume).subscribe(
        _ => {
          this.refreshItems();
          observer.next(true);
          observer.complete();
        },
        error => {
          if (error === CONFLICT) {
            this.popupService.createMsg(STRINGS.EXIST_ENV);
          } else {
            this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
          }
        }
      );
    });
  }

  onAddClick(): boolean {
    this.modalService.createMod({
      title: 'Add Volume',
      items: [
        {
          name: 'name',
          key: 'name',
        },
        {
          name: 'path',
          key: 'path',
        },
      ],
      button: 'Add',
      callback: this.addCallback.bind(this),
    });
    return false;
  }

  refreshItems(): void {
    this.botService.getVolumes(this.current.id).subscribe(
      vols => {
        this.resItems = vols;
      },
      error => {
        this.popupService.createMsg(`unknown error ${error}`);
      }
    );
  }

  ngOnInit(): void {
    this.route.parent.data.subscribe(d => {
      this.current = d.bot;
      this.refreshItems();
    });
  }
}
