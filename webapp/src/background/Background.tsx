import * as React from "react";
import "./Background.css";

// 坐标
class Crood {
    public x: number;
    public y: number;

    constructor(x: number = 0, y: number = 0) {
        this.x = x;
        this.y = y;
    }

    public setCrood = (x: number, y: number) => {
        this.x = x;
        this.y = y;
    };

    public copy = (): Crood => {
        return new Crood(this.x, this.y);
    };
}

// 流星
class ShootingStar {
    private init: Crood;
    private final: Crood;
    private size: number;
    private speed: number;
    private onDestroy: Function;
    private dur: number;
    private pass: number;
    private prev: Crood;
    private now: Crood;
    
    constructor(init = new Crood, final = new Crood, size = 1, speed = 5, onDestroy = null) {
        this.init = init; // 初始位置
        this.final = final; // 最终位置
        this.size = size; // 大小
        this.speed = speed; // 速度：像素/s

        // 飞行总时间
        this.dur = Math.sqrt(Math.pow(this.final.x-this.init.x, 2) + Math.pow(this.final.y-this.init.y, 2)) * 1000 / this.speed;

        this.pass = 0; // 已过去的时间
        this.prev = this.init.copy(); // 上一帧位置
        this.now = this.init.copy(); // 当前位置
        this.onDestroy = onDestroy;
    }
    
    draw = (ctx, delta) => {
        this.pass += delta;
        this.pass = Math.min(this.pass, this.dur);

        let percent = this.pass / this.dur;

        this.now.setCrood(
          this.init.x + (this.final.x - this.init.x) * percent,
          this.init.y + (this.final.y - this.init.y) * percent
        );

        // canvas
        ctx.strokeStyle = '#fff';
        ctx.lineCap = 'round';
        ctx.lineWidth = this.size;
        ctx.beginPath();
        ctx.moveTo(this.now.x, this.now.y);
        ctx.lineTo(this.prev.x, this.prev.y);
        ctx.stroke();

        this.prev.setCrood(this.now.x, this.now.y);
        if (this.pass === this.dur) {
            this.destroy();
        }
    };
    
    destroy = () => {
        this.onDestroy && this.onDestroy();
    };
}

class MeteorShower {
    private cvs: any;
    public ctx: any;
    private stars: any[];
    private T: any;
    private isStop: boolean;
    private playing: boolean;

    constructor(cvs, ctx) {
        this.cvs = cvs;
        this.ctx = ctx;
        this.stars = [];
        this.isStop = false;
        this.playing = false;
    }

    public createStar  = () => {
        let angle = Math.PI / 3;
        let distance = Math.random() * 800;
        let init = new Crood((Math.random() * 1.2 * this.cvs.width - 50), Math.random() * 300|0);
        let final = new Crood(init.x - distance * Math.cos(angle), init.y + distance * Math.sin(angle));
        let size = Math.random() * 2;
        let speed = Math.random() * 20 + 100;
        let star = new ShootingStar(
          init, final, size, speed,
          ()=>{this.remove(star)}
        );
        return star;
    };

    public remove = (star) => {
        this.stars = this.stars.filter((s)=>{ return s !== star});
    };

    public update = (delta) => {
        if (!this.isStop && this.stars.length < 5) {
            this.stars.push(this.createStar());
        }
        this.stars.forEach((star)=>{
            star.draw(this.ctx, delta);
        });
    };

    public tick = () => {
        if (this.playing) return;
        this.playing = true;

        let now = (new Date()).getTime();
        let last = now;
        let delta;

        let  _tick = ()=>{
            if (this.isStop && this.stars.length === 0) {
                cancelAnimationFrame(this.T);
                this.playing = false;
                return;
            }

            delta = now - last;
            delta = delta > 500 ? 30 : (delta < 16? 16 : delta);
            last = now;
            // console.log(delta);

            this.T = requestAnimationFrame(_tick);

            this.ctx.save();
            this.ctx.globalCompositeOperation = "destination-in";
            this.ctx.fillStyle = 'rgba(0,0,0,0.8)'; // 每一帧用 “半透明” 的背景色清除画布
            this.ctx.fillRect(0, 0, this.cvs.width, this.cvs.height);
            this.ctx.restore();
            this.update(delta);
        };
        _tick();
    };

    public start = () => {
        this.isStop = false;
        this.tick();
    };

    public stop = () => {
        this.isStop = true;
    };
}

const Background = (): JSX.Element => {
    const canvas: any = React.createRef();

    React.useEffect(() => {
        let ctx = canvas.current.getContext('2d');

        let meteorShower = new MeteorShower(canvas.current, ctx);
        meteorShower.start();
    }, []);

    return (
      <div className="wrap">
        <canvas ref={canvas} className="background" />
      </div>
    )
};

export default Background;
