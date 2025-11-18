import { defineStore } from 'pinia'
import { ref } from 'vue'
import { request } from '../api/request'
import { API_ENDPOINTS } from '../api/config'

export const useAssetStore = defineStore('asset', () => {
  // 状态
  const assets = ref([])
  const currentAsset = ref(null)
  const jintanAssets = ref([])
  const loading = ref(false)

  // 获取资产列表
  const fetchAssets = async (params = {}) => {
    loading.value = true
    try {
      const response = await request.get(API_ENDPOINTS.ASSET.LIST, params)
      assets.value = response.list || response
      return response
    } catch (error) {
      throw error
    } finally {
      loading.value = false
    }
  }

  // 获取资产详情
  const fetchAssetDetail = async (id) => {
    loading.value = true
    try {
      const response = await request.get(API_ENDPOINTS.ASSET.GET_DETAIL(id))
      currentAsset.value = response
      return response
    } catch (error) {
      throw error
    } finally {
      loading.value = false
    }
  }

  // 获取资产实例列表
  const fetchAssetInstances = async (assetId) => {
    try {
      const response = await request.get(API_ENDPOINTS.ASSET.GET_INSTANCES(assetId))
      return response
    } catch (error) {
      throw error
    }
  }

  // 创建资产
  const createAsset = async (data) => {
    try {
      const response = await request.post(API_ENDPOINTS.ASSET.CREATE, data)
      assets.value.push(response)
      return response
    } catch (error) {
      throw error
    }
  }

  // 更新资产
  const updateAsset = async (id, data) => {
    try {
      const response = await request.put(API_ENDPOINTS.ASSET.UPDATE(id), data)
      const index = assets.value.findIndex(a => a.id === id)
      if (index > -1) {
        assets.value[index] = response
      }
      if (currentAsset.value?.id === id) {
        currentAsset.value = response
      }
      return response
    } catch (error) {
      throw error
    }
  }

  // 删除资产
  const deleteAsset = async (id) => {
    try {
      await request.delete(API_ENDPOINTS.ASSET.DELETE(id))
      assets.value = assets.value.filter(a => a.id !== id)
      if (currentAsset.value?.id === id) {
        currentAsset.value = null
      }
    } catch (error) {
      throw error
    }
  }

  // 获取鲸探资产
  const fetchJintanAssets = async () => {
    loading.value = true
    try {
      const response = await request.get(API_ENDPOINTS.ASSET.GET_JINGTAN)
      jintanAssets.value = response.list || response
      return response
    } catch (error) {
      throw error
    } finally {
      loading.value = false
    }
  }

  // 同步鲸探资产
  const syncJintanAssets = async () => {
    try {
      const response = await request.post(API_ENDPOINTS.ASSET.SYNC_JINGTAN)
      jintanAssets.value = response.list || response
      return response
    } catch (error) {
      throw error
    }
  }

  return {
    assets,
    currentAsset,
    jintanAssets,
    loading,
    fetchAssets,
    fetchAssetDetail,
    fetchAssetInstances,
    createAsset,
    updateAsset,
    deleteAsset,
    fetchJintanAssets,
    syncJintanAssets
  }
})
