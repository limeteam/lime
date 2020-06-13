<template>
    <div class="cate">
        <mt-header :title="title">
            <router-link to="/" slot="left">
                <mt-button icon="back"></mt-button>
            </router-link>
        </mt-header>
        <ul class=vertical-list>
            <li v-for="item in bookList">
                <router-link :to="{path:'book',query:{bookId:item.id}}">
                    <div class="image">
                        <img :src="item.cover" :alt="item.name">
                    </div>
                    <div class="base">
                        <h4>{{item.name}}</h4>
                        <p>{{item.desc}}</p>
                        <div class="author">
                            <i class="icon icon-author"></i>
                            <span>{{item.author}}</span>
                        </div>
                        <div class="category">
                            <span>{{item.type}}</span>
                            <span>{{item.serialize}}</span>
                            <span>{{item.text_num}}万字</span>
                        </div>
                    </div>
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
    import {getBookType} from "../api";

    export default {
        name: "Category",
        data() {
            return {
                listQuery: {
                    page: 1,
                    skip: 0,
                    limit: 20,
                    type: 1
                },
                title: '',
                bookList: []
            }
        },
        methods: {},
        created() {
            this.$store.dispatch('getType', this.$route.query.type).then(res => {
                this.title = res;
                this.listQuery.type = this.$route.query.type;
                getBookType(this.listQuery).then(response => {
                    this.bookList = response.data.result;
                })
            });
        }
    }
</script>
