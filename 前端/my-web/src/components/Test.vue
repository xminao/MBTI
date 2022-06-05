<template>
<div class="start" v-if="start == false && finish == false">
    <el-container class="starttitle">
        <div id="titletext">
        <span><font face="微软雅黑" size="10" color="#fff">准备好了吗？点击开始答题！</font></span>
        </div>
    </el-container>
    <el-container class="startbutton">
        <el-button
        id="startbutton"
        color="#cc9933"
        type="info"
        @click="start = true;select('')">
            开始答题
        </el-button>
    </el-container>
</div>

<div class="test" v-if="start == true && finish == false">
    <el-container class="step">
        <el-progress :percentage="per" :stroke-width="20" :format="format" />
  </el-container>
  <el-container class="title">
      <span><br/><br/><font face="微软雅黑" size="6" color="#696969">{{question_form.desc}}</font></span>
  </el-container>
  <el-container class="options">
    <el-radio-group v-model="option" size="large">
        <button class="button"
            label="A" 
            @click="select('A')"
            >{{question_form.option_a}}</button>
    </el-radio-group>
  </el-container>
  <el-container class="options">
    <el-radio-group v-model="option" size="large">
        <button class="button"
            label="B" 
            @click="select('A')"
            >{{question_form.option_b}}</button>
    </el-radio-group>
  </el-container>
</div>

<div class="end" v-if="finish == true">
    <el-container class="endtitle">
        <div id="endtext">
        <span><font face="微软雅黑" size="10" color="#fff">完成答题啦</font></span>
        </div>
    </el-container>
        
    <el-container class="endcontent">
        {{result_form}} {{result}}
    </el-container>
</div>

</template>

<script>
import { reactive, ref, getCurrentInstance } from 'vue'

export default ({
    setup() {
        const active = ref(0)
        const button_select = ref(false)
        const start = ref(false)
        const index = ref(0)
        const per = ref(0)
        let finish = ref(false)
        let result = ref('')

        const {proxy} = getCurrentInstance()

        let result_form = reactive({
            E: 0,
            I: 0,
            S: 0,
            N: 0,
            T: 0,
            F: 0,
            J: 0,
            P: 0,
        })

        let getResult=()=> {
            if (result_form.E > result_form.I) {
                result = result + 'E'
            } else {
                result = result + 'I'
            }

            if (result_form.S > result_form.N) {
                result = result + 'S'
            } else {
                result = result + 'N'
            }

            if (result_form.T > result_form.F) {
                result = result + 'T'
            } else {
                result = result + 'F'
            }

            if (result_form.J > result_form.P) {
                result = result + 'J'
            } else {
                result = result + 'P'
            }
        }

        let question_form = reactive({
            id: '',
            desc: '',
            option_a: '',
            option_b: '',
        })

        const getidList=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            return listres
        }

        const select=async(option)=> {
            if (option != '') {
                index.value++
            }
            const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            const idList = listres.question_id_list
            const num = idList.length
            if (index.value >= num) {
                per.value = (index.value/num)*100
                getResult()
                console.log(result)
                finish.value = true
                return
            }
            const obj = {"id":idList[index.value]}
            const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
            question_form.desc = res.question_info.question_desc
            question_form.option_a = res.question_info.option_a_desc
            question_form.option_b = res.question_info.option_b_desc
            let target
            if (option == 'A') {
                target = res.question_info.option_a_target
            } else if (option == 'B') {
                target = res.question_info.option_b_target
            }

            switch (target) {
                case 'E':
                    result_form.E++
                    break;
                case 'I':
                    result_form.I++
                    break;
                case 'S':
                    result_form.S++
                    break;
                case 'N':
                    result_form.N++
                    break;
                case 'T':
                    result_form.T++
                    break;
                case 'F':
                    result_form.F++
                    break;
                case 'J':
                    result_form.J++
                    break;
                case 'P':
                    result_form.P++
                    break;
                default:
                    break;
            }
            //console.log(result_form)
           // index.value++
            per.value = (index.value/listres.question_id_list.length)*100
        }

        const format = (percentage) => (percentage === 100 ? '完成' : `${percentage}%`)

        return {
            select,
            per,
            index,
            start,
            finish,
            question_form,
            result_form,
            result,
            button_select, 
            active, 
            format}
    },
})
</script>

<style scoped>
.start {
    background-color: #fff;
    height: 100%;
    padding-inline: 0; 
}
.starttitle {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
.starttitle #titletext {
    position: relative;
    top: 90px;
}
.startbutton {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}

.startbutton .el-button {
    position: relative;
    top: 20px;
    color: #000;
    font-size: 20px;
    height: 70px;
    width: 25%;
    border-radius: 50px
}
.test {
    justify-content: center;
    background-color: #fff;
    height: 100%;
    padding-inline: 0;
}
.step {
        justify-content: center;
        background-color: #fff;
        position: relative;
        top: 20px;
        height: 10%;
        padding: 0%;
}

.el-progress--line {
    margin-top: 15px;
    margin-bottom: 15px;
    width: 35%;
}

.el-radio-group {
    justify-content: center;
    width: 100%;
}

#startbutton {
    color: #fff;
    font-size: 30px;
    height: 70px;
    width: 20%;
}

.button {
    background-color: #9999CC; /* Green */
    border: #fff;
    border-radius: 12px;
    color: white;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    height: 70px;
    width: 35%;
    font-size: 20px;
}

.title {
    justify-content: center;
    background-color: #fff;
    height: 30%;
    position: relative;
    top: 15px;
    padding: 0%; 
}

.options {
    justify-content: center;
    background-color: #fff;
    height: 20%;
    width: 100%;
    padding: 0%;
}

.end {
    background-color: #fff;
    height: 100%;
    padding-inline: 0; 
}

.endtitle {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
.endtitle #endtext {
    position: relative;
    top: 90px;
}
.endcontent {
    justify-content: center;
    height: 50%;
    background-color: #4298B4;
    width: 100%;
}
</style>
