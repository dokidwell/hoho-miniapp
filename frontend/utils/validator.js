/**
 * 表单验证工具
 * 提供常用的验证规则和方法
 */

class Validator {
  /**
   * 验证手机号
   * @param {string} phone - 手机号
   * @returns {boolean} 是否有效
   */
  static isValidPhone(phone) {
    if (!phone) return false
    const phoneReg = /^1[3-9]\d{9}$/
    return phoneReg.test(phone)
  }

  /**
   * 验证邮箱
   * @param {string} email - 邮箱
   * @returns {boolean} 是否有效
   */
  static isValidEmail(email) {
    if (!email) return false
    const emailReg = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/
    return emailReg.test(email)
  }

  /**
   * 验证密码强度
   * @param {string} password - 密码
   * @returns {Object} 验证结果
   */
  static validatePassword(password) {
    if (!password) {
      return {
        valid: false,
        message: '密码不能为空',
        strength: 0
      }
    }

    if (password.length < 6) {
      return {
        valid: false,
        message: '密码长度至少6位',
        strength: 0
      }
    }

    if (password.length > 20) {
      return {
        valid: false,
        message: '密码长度不能超过20位',
        strength: 0
      }
    }

    // 计算密码强度
    let strength = 0
    
    // 包含小写字母
    if (/[a-z]/.test(password)) strength++
    // 包含大写字母
    if (/[A-Z]/.test(password)) strength++
    // 包含数字
    if (/\d/.test(password)) strength++
    // 包含特殊字符
    if (/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password)) strength++
    // 长度大于10
    if (password.length > 10) strength++

    return {
      valid: true,
      strength: Math.min(strength, 5),
      message: this.getPasswordStrengthText(strength)
    }
  }

  /**
   * 获取密码强度文本
   * @param {number} strength - 强度值
   * @returns {string} 强度文本
   */
  static getPasswordStrengthText(strength) {
    const texts = ['很弱', '弱', '一般', '强', '很强']
    return texts[Math.min(strength - 1, 4)] || '很弱'
  }

  /**
   * 验证身份证号
   * @param {string} idCard - 身份证号
   * @returns {boolean} 是否有效
   */
  static isValidIdCard(idCard) {
    if (!idCard) return false
    
    // 18位身份证号
    const idCardReg = /^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]$/
    
    if (!idCardReg.test(idCard)) {
      return false
    }

    // 验证校验码
    const factors = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
    const checkCodes = ['1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2']
    
    let sum = 0
    for (let i = 0; i < 17; i++) {
      sum += parseInt(idCard[i]) * factors[i]
    }
    
    const checkCode = checkCodes[sum % 11]
    return checkCode === idCard[17].toUpperCase()
  }

  /**
   * 验证URL
   * @param {string} url - URL
   * @returns {boolean} 是否有效
   */
  static isValidUrl(url) {
    if (!url) return false
    const urlReg = /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
    return urlReg.test(url)
  }

  /**
   * 验证数字
   * @param {*} value - 值
   * @param {Object} options - 选项
   * @returns {Object} 验证结果
   */
  static validateNumber(value, options = {}) {
    const {
      min = -Infinity,
      max = Infinity,
      integer = false,
      positive = false
    } = options

    const num = Number(value)

    if (isNaN(num)) {
      return {
        valid: false,
        message: '请输入有效的数字'
      }
    }

    if (integer && !Number.isInteger(num)) {
      return {
        valid: false,
        message: '请输入整数'
      }
    }

    if (positive && num <= 0) {
      return {
        valid: false,
        message: '请输入正数'
      }
    }

    if (num < min) {
      return {
        valid: false,
        message: `数值不能小于${min}`
      }
    }

    if (num > max) {
      return {
        valid: false,
        message: `数值不能大于${max}`
      }
    }

    return {
      valid: true,
      value: num
    }
  }

  /**
   * 验证字符串长度
   * @param {string} str - 字符串
   * @param {Object} options - 选项
   * @returns {Object} 验证结果
   */
  static validateLength(str, options = {}) {
    const {
      min = 0,
      max = Infinity,
      trim = true
    } = options

    const value = trim ? (str || '').trim() : (str || '')
    const length = value.length

    if (length < min) {
      return {
        valid: false,
        message: `长度不能少于${min}个字符`
      }
    }

    if (length > max) {
      return {
        valid: false,
        message: `长度不能超过${max}个字符`
      }
    }

    return {
      valid: true,
      value: value
    }
  }

  /**
   * 验证表单
   * @param {Object} data - 表单数据
   * @param {Object} rules - 验证规则
   * @returns {Object} 验证结果
   */
  static validateForm(data, rules) {
    const errors = {}

    for (const field in rules) {
      const fieldRules = rules[field]
      const value = data[field]

      for (const rule of fieldRules) {
        const { type, message, validator, ...options } = rule

        let isValid = true
        let errorMessage = message

        switch (type) {
          case 'required':
            isValid = value !== undefined && value !== null && value !== ''
            errorMessage = message || `${field}不能为空`
            break

          case 'phone':
            isValid = this.isValidPhone(value)
            errorMessage = message || '请输入有效的手机号'
            break

          case 'email':
            isValid = this.isValidEmail(value)
            errorMessage = message || '请输入有效的邮箱'
            break

          case 'password':
            const passwordResult = this.validatePassword(value)
            isValid = passwordResult.valid
            errorMessage = passwordResult.message
            break

          case 'idCard':
            isValid = this.isValidIdCard(value)
            errorMessage = message || '请输入有效的身份证号'
            break

          case 'url':
            isValid = this.isValidUrl(value)
            errorMessage = message || '请输入有效的URL'
            break

          case 'number':
            const numberResult = this.validateNumber(value, options)
            isValid = numberResult.valid
            errorMessage = numberResult.message
            break

          case 'length':
            const lengthResult = this.validateLength(value, options)
            isValid = lengthResult.valid
            errorMessage = lengthResult.message
            break

          case 'custom':
            if (validator && typeof validator === 'function') {
              const result = validator(value, data)
              isValid = result === true || result.valid === true
              errorMessage = result.message || message
            }
            break

          default:
            break
        }

        if (!isValid) {
          errors[field] = errorMessage
          break // 一个字段只显示第一个错误
        }
      }
    }

    return {
      valid: Object.keys(errors).length === 0,
      errors: errors
    }
  }

  /**
   * 格式化手机号
   * @param {string} phone - 手机号
   * @returns {string} 格式化后的手机号
   */
  static formatPhone(phone) {
    if (!phone) return ''
    const cleaned = phone.replace(/\D/g, '')
    const match = cleaned.match(/^(\d{3})(\d{4})(\d{4})$/)
    if (match) {
      return `${match[1]} ${match[2]} ${match[3]}`
    }
    return phone
  }

  /**
   * 脱敏手机号
   * @param {string} phone - 手机号
   * @returns {string} 脱敏后的手机号
   */
  static maskPhone(phone) {
    if (!phone || phone.length !== 11) return phone
    return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
  }

  /**
   * 脱敏身份证号
   * @param {string} idCard - 身份证号
   * @returns {string} 脱敏后的身份证号
   */
  static maskIdCard(idCard) {
    if (!idCard || idCard.length !== 18) return idCard
    return idCard.replace(/(\d{6})\d{8}(\d{4})/, '$1********$2')
  }

  /**
   * 脱敏姓名
   * @param {string} name - 姓名
   * @returns {string} 脱敏后的姓名
   */
  static maskName(name) {
    if (!name) return ''
    if (name.length <= 1) return name
    if (name.length === 2) return name[0] + '*'
    return name[0] + '*'.repeat(name.length - 2) + name[name.length - 1]
  }
}

export default Validator
