<template>
  <div class="timeline">
    <div v-for="(process, index) in processes" :key="index" class="timeline-item">
      <div class="timeline-circle"></div>
      <div class="timeline-content">
        <h3>{{ process.Step }}</h3>
        <p>{{ process.Performer }}</p>
        <span>{{ process.Timestamp }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Timeline',
  props: {
    processes: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style scoped>
.timeline {
  position: relative;
  padding-left: 40px; /* 增加间距以适应更大的圆圈 */
}

/* 添加时间轴线条样式 */
.timeline::before {
  content: '';
  position: absolute;
  top: 0;
  left: 24px; /* 调整位置以匹配圆圈半径 */
  height: 100%;
  width: 6px; /* 增加线条宽度 */
  background: linear-gradient(to bottom, #34a853, #005f4f); /* 更新颜色渐变 */
  z-index: -1;
}

.timeline-item {
  position: relative;
  margin-bottom: 60px; /* 增加间距 */
}

/* 使用伪元素替代真实的 div.timeline-line */
.timeline-item::after {
  content: '';
  position: absolute;
  top: 24px; /* 与圆圈顶部对齐 */
  left: 20px; /* 略微调整以适应新宽度 */
  width: 6px; /* 与主线条宽度一致 */
  height: 20px; /* 调整为与圆圈高度匹配的小段 */
  background-color: white; /* 隐藏线条 */
  z-index: 1;
}

.timeline-circle {
  position: absolute;
  left: 0;
  top: 24px; /* 调整以适应新的线条和间距 */
  width: 48px; /* 增加大小 */
  height: 48px;
  background-color: #34a853; /* 更新背景色 */
  border: 4px solid white;
  border-radius: 50%;
  box-shadow: 0 0 0 6px rgba(52, 168, 83, 0.3); /* 更新阴影 */
  transition: transform 0.3s ease, background-color 0.3s ease; /* 添加动画效果 */
}

.timeline-circle:hover {
  transform: scale(1.1); /* 悬停时放大 */
  background-color: #005f4f; /* 更新悬停颜色 */
}

.timeline-content {
  padding-left: 70px; /* 调整以适应更大的圆圈 */
}

.timeline-content h3 {
  margin: 0;
  font-size: 1.75em; /* 增加字体大小 */
  font-weight: bold;
  color: #333; /* 更新字体颜色以提高可读性 */
}

.timeline-content p {
  margin: 10px 0;
  color: #666; /* 更新字体颜色 */
}

.timeline-content span {
  color: #999; /* 更新时间戳颜色 */
  font-size: 0.95em;
}
</style>
